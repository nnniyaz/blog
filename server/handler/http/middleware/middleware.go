package middleware

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nnniyaz/blog/handler/http/response"
	"github.com/nnniyaz/blog/pkg/logger"
	"github.com/nnniyaz/blog/pkg/web"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
)

type Middleware struct {
	logger logger.Logger
	client mongo.Client
}

func New(l logger.Logger, client *mongo.Client) *Middleware {
	return &Middleware{logger: l, client: *client}
}

func (m *Middleware) Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if e, ok := err.(error); ok {
					response.NewInternal(m.logger, w, r, e, zap.Any("panic", err))
				} else {
					response.NewInternal(m.logger, w, r, nil, zap.Any("panic", err))
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) Trace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceId := web.GenerateTraceId()
		ctx := context.WithValue(r.Context(), "traceId", traceId)
		ctx = context.WithValue(ctx, middleware.RequestIDKey, traceId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) RequestInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "requestInfo", web.GetRequestInfo(r))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) PaginationParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limit, err := getParamInt64(r, "limit")
		if err != nil {
			limit = 25
		}

		offset, err := getParamInt64(r, "offset")
		if err != nil {
			offset = 0
		}

		isDeleted, err := getParamBool(r, "is_deleted")
		if err != nil {
			isDeleted = false
		}

		search := getParam(r, "search")

		ctx := context.WithValue(r.Context(), "limit", limit)
		ctx = context.WithValue(ctx, "offset", offset)
		ctx = context.WithValue(ctx, "is_deleted", isDeleted)
		ctx = context.WithValue(ctx, "search", search)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) WithTransaction(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch || r.Method == http.MethodDelete {
			session, err := m.client.StartSession()
			if err != nil {
				response.NewError(m.logger, w, r, err)
				return
			}
			defer session.EndSession(r.Context())

			_, err = session.WithTransaction(r.Context(), func(sessCtx mongo.SessionContext) (interface{}, error) {
				ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

				next.ServeHTTP(ww, r.WithContext(sessCtx))

				if ww.Status() >= 500 {
					return nil, fmt.Errorf("status code %d", ww.Status())
				}

				return nil, nil
			})
			if err != nil {
				m.logger.Error("error while executing transaction", zap.Error(err))
				return
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
