package http

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/nnniyaz/blog/internal/handlers/http/article"
	authHandler "github.com/nnniyaz/blog/internal/handlers/http/auth"
	"github.com/nnniyaz/blog/internal/handlers/http/author"
	"github.com/nnniyaz/blog/internal/handlers/http/bio"
	"github.com/nnniyaz/blog/internal/handlers/http/book"
	"github.com/nnniyaz/blog/internal/handlers/http/contact"
	"github.com/nnniyaz/blog/internal/handlers/http/middleware"
	"github.com/nnniyaz/blog/internal/handlers/http/project"
	userHandler "github.com/nnniyaz/blog/internal/handlers/http/user"
	"github.com/nnniyaz/blog/internal/services"
	"github.com/nnniyaz/blog/pkg/logger"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Handler struct {
	middleware *middleware.Middleware
	article    *articleHandler.HttpDelivery
	contact    *contactHandler.HttpDelivery
	author     *authorHandler.HttpDelivery
	bio        *bioHandler.HttpDelivery
	book       *bookHandler.HttpDelivery
	project    *projectHandler.HttpDelivery
	user       *userHandler.HttpDelivery
	auth       *authHandler.HttpDelivery
}

func NewHandler(l logger.Logger, client *mongo.Client, services *services.Service) *Handler {
	return &Handler{
		middleware: middleware.New(l, client),
		article:    articleHandler.NewHttpDelivery(l, services.Article),
		contact:    contactHandler.NewHttpDelivery(l, services.Contact),
		author:     authorHandler.NewHttpDelivery(l, services.Author),
		bio:        bioHandler.NewHttpDelivery(l, services.Bio),
		book:       bookHandler.NewHttpDelivery(l, services.Book),
		project:    projectHandler.NewHttpDelivery(l, services.Project),
		user:       userHandler.NewHttpDelivery(l, services.User),
		auth:       authHandler.NewHttpDelivery(l, services.Auth),
	}
}

func (h *Handler) InitRoutes(isDevMode bool) *chi.Mux {
	r := chi.NewRouter()

	if isDevMode {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins: []string{
				"http://localhost:3000",
				"https://localhost:3000",
				"http://localhost:3001",
				"https://localhost:3001",
			},
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
		}))
	} else {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins: []string{
				"",
				"",
			},
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
		}))
	}

	r.Use(h.middleware.Recover)
	r.Use(h.middleware.Trace)
	r.Use(h.middleware.RequestInfo)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.RealIP)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/auth", func(r chi.Router) {
		r.With(h.middleware.NoAuth).Post("/login", h.auth.Login)
		r.With(h.middleware.UserAuth).Post("/logout", h.auth.Logout)
	})

	r.Route("/about", func(r chi.Router) {
		r.Route("/author", func(r chi.Router) {
			r.Get("/{id}", h.author.GetAuthor)

			r.Use(h.middleware.UserAuth)

			r.Post("/", h.author.CreateAuthor)
			r.Put("/{id}", h.author.UpdateAuthor)
			r.Delete("/{id}", h.author.DeleteAuthor)
			r.Put("/restore/{id}", h.author.RestoreAuthor)
			r.With(h.middleware.PaginationParams).Get("/", h.author.GetAllAuthors)
		})

		r.Route("/bio", func(r chi.Router) {
			r.Get("/{id}", h.bio.GetBio)

			r.Use(h.middleware.UserAuth)

			r.Post("/", h.bio.CreateBio)
			r.Put("/{id}", h.bio.UpdateBio)
			r.Delete("/{id}", h.bio.DeleteBio)
			r.Put("/restore/{id}", h.bio.RestoreBio)
			r.Get("/active", h.bio.GetActiveBio)
			r.Get("/", h.bio.GetAllBios)
		})

		r.Route("/contact", func(r chi.Router) {
			r.With(h.middleware.PaginationParams).Get("/", h.contact.GetAllContacts)

			r.Use(h.middleware.UserAuth)

			r.Post("/", h.contact.CreateContact)
			r.Put("/{id}", h.contact.UpdateContact)
			r.Delete("/{id}", h.contact.DeleteContact)
			r.Put("/restore/{id}", h.contact.RestoreContact)
			r.Get("/{id}", h.contact.GetContact)
		})
	})

	r.Route("/project", func(r chi.Router) {
		r.With(h.middleware.PaginationParams).Get("/", h.project.GetAllProjects)

		r.Use(h.middleware.UserAuth)

		r.Post("/", h.project.CreateProject)
		r.Put("/{id}", h.project.UpdateProject)
		r.Delete("/{id}", h.project.DeleteProject)
		r.Put("/restore/{id}", h.project.RestoreProject)
		r.Get("/{id}", h.project.GetProject)
	})

	r.Route("/article", func(r chi.Router) {
		r.With(h.middleware.PaginationParams).Get("/", h.article.GetAllArticles)

		r.Use(h.middleware.UserAuth)

		r.Post("/", h.article.CreateArticle)
		r.Put("/{id}", h.article.UpdateArticle)
		r.Delete("/{id}", h.article.DeleteArticle)
		r.Put("/restore/{id}", h.article.RestoreArticle)
		r.Get("/{id}", h.article.GetArticle)
	})

	r.Route("/book", func(r chi.Router) {
		r.With(h.middleware.PaginationParams).Get("/", h.book.GetAllBooks)

		r.Use(h.middleware.UserAuth)

		r.Post("/", h.book.CreateBook)
		r.Put("/{id}", h.book.UpdateBook)
		r.Delete("/{id}", h.book.DeleteBook)
		r.Put("/restore/{id}", h.book.RestoreBook)
		r.Get("/{id}", h.book.GetBook)
	})

	r.Route("/user", func(r chi.Router) {
		r.Use(h.middleware.UserAuth)

		r.Put("/", h.user.GetAllUsers)
		r.Put("/{id}", h.user.GetUser)

		r.Use(h.middleware.AdminCheck)

		r.Post("/", h.user.CreateUser)
		r.Put("/email/{id}", h.user.UpdateUserEmail)
		r.Put("/password/{id}", h.user.UpdateUserPassword)
		r.Put("/role/{id}", h.user.UpdateUserRole)
		r.Delete("/{id}", h.user.DeleteUser)
		r.Put("/restore/{id}", h.user.RestoreUser)
	})

	return r
}
