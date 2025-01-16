package http

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/nnniyaz/blog/internal/handlers/http/article"
	"github.com/nnniyaz/blog/internal/handlers/http/author"
	"github.com/nnniyaz/blog/internal/handlers/http/bio"
	"github.com/nnniyaz/blog/internal/handlers/http/book"
	"github.com/nnniyaz/blog/internal/handlers/http/contact"
	"github.com/nnniyaz/blog/internal/handlers/http/middleware"
	"github.com/nnniyaz/blog/internal/handlers/http/project"
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

	r.Route("/about", func(r chi.Router) {
		r.Route("/author", func(r chi.Router) {
			r.With(h.middleware.PaginationParams).Get("/", h.author.GetAllAuthors)
			r.Post("/", h.author.CreateAuthor)
			r.Get("/{id}", h.author.GetAuthor)
			r.Put("/{id}", h.author.UpdateAuthor)
			r.Delete("/{id}", h.author.DeleteAuthor)
			r.Put("/restore/{id}", h.author.RestoreAuthor)
		})

		r.Route("/bio", func(r chi.Router) {
			r.Get("/", h.bio.GetAllBios)
			r.Post("/", h.bio.CreateBio)
			r.Get("/{id}", h.bio.GetBio)
			r.Get("/active", h.bio.GetActiveBio)
			r.Put("/{id}", h.bio.UpdateBio)
			r.Delete("/{id}", h.bio.DeleteBio)
			r.Put("/restore/{id}", h.bio.RestoreBio)
		})

		r.Route("/contact", func(r chi.Router) {
			r.With(h.middleware.PaginationParams).Get("/", h.contact.GetAllContacts)
			r.Post("/", h.contact.CreateContact)
			r.Get("/{id}", h.contact.GetContact)
			r.Put("/{id}", h.contact.UpdateContact)
			r.Delete("/{id}", h.contact.DeleteContact)
			r.Put("/restore/{id}", h.contact.RestoreContact)
		})
	})

	r.Route("/project", func(r chi.Router) {
		r.With(h.middleware.PaginationParams).Get("/", h.project.GetAllProjects)
		r.Post("/", h.project.CreateProject)
		r.Get("/{id}", h.project.GetProject)
		r.Put("/{id}", h.project.UpdateProject)
		r.Delete("/{id}", h.project.DeleteProject)
		r.Put("/restore/{id}", h.project.RestoreProject)
	})

	r.Route("/article", func(r chi.Router) {
		r.With(h.middleware.PaginationParams).Get("/", h.article.GetAllArticles)
		r.Post("/", h.article.CreateArticle)
		r.Get("/{id}", h.article.GetArticle)
		r.Put("/{id}", h.article.UpdateArticle)
		r.Delete("/{id}", h.article.DeleteArticle)
		r.Put("/restore/{id}", h.article.RestoreArticle)
	})

	r.Route("/book", func(r chi.Router) {
		r.With(h.middleware.PaginationParams).Get("/", h.book.GetAllBooks)
		r.Post("/", h.book.CreateBook)
		r.Get("/{id}", h.book.GetBook)
		r.Put("/{id}", h.book.UpdateBook)
		r.Delete("/{id}", h.book.DeleteBook)
		r.Put("/restore/{id}", h.book.RestoreBook)
	})

	return r
}
