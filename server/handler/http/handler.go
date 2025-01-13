package http

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	articleHandler "github.com/nnniyaz/blog/handler/http/article"
	authorHandler "github.com/nnniyaz/blog/handler/http/author"
	bioHandler "github.com/nnniyaz/blog/handler/http/bio"
	bookHandler "github.com/nnniyaz/blog/handler/http/book"
	contactHandler "github.com/nnniyaz/blog/handler/http/contact"
	"github.com/nnniyaz/blog/handler/http/middleware"
	"github.com/nnniyaz/blog/pkg/logger"
	"github.com/nnniyaz/blog/service"
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
}

func NewHandler(l logger.Logger, client *mongo.Client, services *service.Service) *Handler {
	return &Handler{
		middleware: middleware.New(l, client),
		article:    articleHandler.NewHttpDelivery(l, services.Article),
		contact:    contactHandler.NewHttpDelivery(l, services.Contact),
		author:     authorHandler.NewHttpDelivery(l, services.Author),
		bio:        bioHandler.NewHttpDelivery(l, services.Bio),
		book:       bookHandler.NewHttpDelivery(l, services.Book),
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
			r.Get("/", h.author.GetAllAuthors)
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
			r.Put("/{id}", h.bio.UpdateBio)
			r.Delete("/{id}", h.bio.DeleteBio)
			r.Put("/restore/{id}", h.bio.RestoreBio)
		})

		r.Route("/contact", func(r chi.Router) {
			r.Get("/", h.contact.GetAllContacts)
			r.Post("/", h.contact.CreateContact)
			r.Get("/{id}", h.contact.GetContact)
			r.Put("/{id}", h.contact.UpdateContact)
			r.Delete("/{id}", h.contact.DeleteContact)
			r.Put("/restore/{id}", h.contact.RestoreContact)
		})
	})

	r.Route("/project", func(r chi.Router) {
		//r.Get("/", h.GetProjects)
		//r.Post("/", h.CreateProject)
		//r.Get("/{id}", h.GetProject)
		//r.Put("/{id}", h.UpdateProject)
		//r.Delete("/{id}", h.DeleteProject)
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
		r.Get("/", h.book.GetAllBooks)
		r.Post("/", h.book.CreateBook)
		r.Get("/{id}", h.book.GetBook)
		r.Put("/{id}", h.book.UpdateBook)
		r.Delete("/{id}", h.book.DeleteBook)
		r.Put("/restore/{id}", h.book.RestoreBook)
	})

	return r
}
