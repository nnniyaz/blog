package http

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	articleHandler "github.com/nnniyaz/blog/handler/http/article"
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
}

func NewHandler(l logger.Logger, client *mongo.Client, services *service.Service) *Handler {
	return &Handler{
		middleware: middleware.New(l, client),
		article:    articleHandler.NewHttpDelivery(l, services.Article),
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
		r.Route("/me", func(r chi.Router) {
			//r.Get("/", h.GetMe)
			//r.Post("/", h.CreateMe)
			//r.Get("/{id}", h.GetMe)
			//r.Put("/{id}", h.UpdateMe)
			//r.Delete("/{id}", h.DeleteMe)
		})

		r.Route("/bio", func(r chi.Router) {
			//r.Get("/", h.GetBio)
			//r.Post("/", h.CreateBio)
			//r.Get("/{id}", h.GetBio)
			//r.Put("/{id}", h.UpdateBio)
			//r.Delete("/{id}", h.DeleteBio)
		})

		r.Route("/contact", func(r chi.Router) {
			//r.Get("/", h.GetContacts)
			//r.Post("/", h.CreateContact)
			//r.Get("/{id}", h.GetContact)
			//r.Put("/{id}", h.UpdateContact)
			//r.Delete("/{id}", h.DeleteContact)
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
		r.With(h.middleware.PaginationParams).Get("/", h.article.FindAllArticles)
		r.Post("/", h.article.CreateArticle)
		r.Get("/{id}", h.article.FindArticle)
		r.Put("/{id}", h.article.UpdateArticle)
		r.Delete("/{id}", h.article.DeleteArticle)
		r.Put("/restore/{id}", h.article.RestoreArticle)
	})

	r.Route("/book", func(r chi.Router) {
		//r.Get("/", h.GetBooks)
		//r.Post("/", h.CreateBook)
		//r.Get("/{id}", h.GetBook)
		//r.Put("/{id}", h.UpdateBook)
		//r.Delete("/{id}", h.DeleteBook)
	})

	return r
}
