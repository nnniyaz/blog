package main

import (
	"context"
	"errors"
	server "github.com/nnniyaz/blog"
	"github.com/nnniyaz/blog/domain/base/config"
	hHttp "github.com/nnniyaz/blog/handler/http"
	"github.com/nnniyaz/blog/pkg/email"
	"github.com/nnniyaz/blog/pkg/env"
	"github.com/nnniyaz/blog/pkg/logger"
	"github.com/nnniyaz/blog/pkg/mongo"
	"github.com/nnniyaz/blog/repo"
	"github.com/nnniyaz/blog/service"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//	@title			Scoring System API
//	@version		0.0.1
//	@description	Detailed info about all endpoints

//	@contact.name	API Support
//	@contact.url	https://t.me/nassyrovich

//	@host		http://localhost:8080
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

	// --- init email service
	email, err := email.New(cfg.GetSmtpHost(), cfg.GetSmtpPort(), cfg.GetSmtpUser(), cfg.GetSmtpPass())
	if err != nil {
		lg.Fatal("failed to init email service", zap.Error(err))
	}

	// --- init handler
	repos := repo.NewRepo(db)
	services := service.NewService(repos, cfg, lg, email)
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
