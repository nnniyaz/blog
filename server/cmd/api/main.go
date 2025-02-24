package main

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nnniyaz/blog/internal/domain/base/config"
	hHttp "github.com/nnniyaz/blog/internal/handlers/http"
	"github.com/nnniyaz/blog/internal/repos"
	"github.com/nnniyaz/blog/internal/services"
	"github.com/nnniyaz/blog/pkg/email"
	"github.com/nnniyaz/blog/pkg/env"
	"github.com/nnniyaz/blog/pkg/logger"
	"github.com/nnniyaz/blog/pkg/mongo"
	"github.com/nnniyaz/blog/pkg/server"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//	@title			Personal Blog API
//	@version		0.0.1
//	@description	Detailed info about all endpoints

//	@contact.name	API Support
//	@contact.url	https://t.me/niyaznassyrov

//	@host		https://api.nassyrov.net
//	@schemes	https

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Cookie
func main() {
	start := time.Now()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	ctx = context.WithValue(ctx, "start", start)

	// --- init config
	cfg := config.NewConfig(
		env.MustGetEnvAsInt("PORT"),
		env.MustGetEnvAsBool("IS_DEV_MODE"),
		env.MustGetEnv("MONGO_URI"),
		env.MustGetEnv("SMTP_HOST"),
		env.MustGetEnvAsInt("SMTP_PORT"),
		env.MustGetEnv("SMTP_USER"),
		env.MustGetEnv("SMTP_PASS"),
		env.MustGetEnv("SPACE_BUCKET"),
		env.MustGetEnv("SPACE_KEY"),
		env.MustGetEnv("SPACE_SECRET"),
		env.MustGetEnv("SPACE_ENDPOINT"),
		env.MustGetEnv("SPACE_REGION"),
		env.MustGetEnv("SPACE_NAME"),
		env.MustGetEnv("SPACE_HOST"),
	)

	// --- init logger
	lg, err := logger.NewLogger(cfg.GetIsDevMode())
	if err != nil {
		panic(err)
	}
	defer lg.Sync()

	// --- init mongo db
	db, err := mongo.New(ctx, cfg.GetMongoUri())
	if err != nil {
		lg.Fatal("failed to init mongodb", zap.Error(err))
		return
	}

	// --- init s3 client
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(cfg.GetSpaceKey(), cfg.GetSpaceSecret(), ""),
		Endpoint:    aws.String(cfg.GetSpaceEndPoint()),
		Region:      aws.String(cfg.GetSpaceRegion()),
	}
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		lg.Fatal("failed to init s3 session", zap.Error(err))
		return
	}
	s3Client := s3.New(newSession)

	// --- init email service
	email, err := email.New(cfg.GetSmtpHost(), cfg.GetSmtpPort(), cfg.GetSmtpUser(), cfg.GetSmtpPass())
	if err != nil {
		lg.Fatal("failed to init email services", zap.Error(err))
	}

	// --- init handler
	repos := repos.NewRepo(db)
	services := services.NewService(repos, cfg, s3Client, email)
	handlers := hHttp.NewHandler(lg, db, services)

	// --- init server
	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg.GetPort(), handlers.InitRoutes(cfg.GetIsDevMode()), start); err != nil && errors.Is(http.ErrServerClosed, err) {
			lg.Fatal("error occurred while running http server: ", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		lg.Fatal("error occurred on server shutting down: ", zap.Error(err))
	}
}
