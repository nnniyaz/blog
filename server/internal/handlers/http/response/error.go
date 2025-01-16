package response

import (
	"context"
	"github.com/nnniyaz/blog/pkg/core"
	"github.com/nnniyaz/blog/pkg/logger"
	"go.uber.org/zap"
	"net/http"
)

var (
	ErrUnauthorized = core.NewI18NError(core.EUNAUTHORIZED, core.TXT_UNAUTHORIZED)
)

type Error struct {
	TraceId  string   `json:"traceId" example:"string"`
	Success  bool     `json:"success" example:"false"`
	Messages []string `json:"messages" example:"error message"`
}

func NewUnauthorized(l logger.Logger, w http.ResponseWriter, r *http.Request) {
	errMsg := core.NewI18NError(core.EUNAUTHORIZED, core.TXT_UNAUTHORIZED).ErrMsg()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	writeErr := writeJSON(w, Error{
		TraceId:  r.Context().Value("traceId").(string),
		Success:  false,
		Messages: []string{errMsg},
	})
	if writeErr != nil {
		l.Error("failed to write json", zap.Error(writeErr))
	}
}

func NewBad(l logger.Logger, w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	writeErr := writeJSON(w, Error{
		TraceId:  r.Context().Value("traceId").(string),
		Success:  false,
		Messages: []string{err.Error()},
	})
	if writeErr != nil {
		l.Error("failed to write json", zap.Error(writeErr))
	}
}

func NewInternal(l logger.Logger, w http.ResponseWriter, r *http.Request, err error, fields ...zap.Field) {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	} else {
		errMsg = core.NewI18NError(core.EINTERNAL, core.TXT_UNKNOWN_ERROR).ErrMsg()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	writeErr := writeJSON(w, Error{
		TraceId:  r.Context().Value("traceId").(string),
		Success:  false,
		Messages: []string{errMsg},
	})
	if writeErr != nil {
		l.Error("failed to write json", zap.Error(writeErr))
	}
	l.Error("internal error", fields...)
}

func NewError(l logger.Logger, w http.ResponseWriter, r *http.Request, err error) {
	if err == context.Canceled || err == context.DeadlineExceeded {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	i18nError, ok := err.(*core.I18NError)
	if ok {
		switch i18nError.Code {
		case core.EINTERNAL:
			w.WriteHeader(http.StatusInternalServerError)
		case core.ENOTFOUND:
			w.WriteHeader(http.StatusNotFound)
		case core.EINVALID:
			w.WriteHeader(http.StatusBadRequest)
		case core.ECONFLICT:
			w.WriteHeader(http.StatusConflict)
		case core.EUNAUTHORIZED:
			w.WriteHeader(http.StatusUnauthorized)
			http.SetCookie(w, &http.Cookie{
				Name:     "ardo-app-session",
				Value:    "",
				Path:     "/",
				HttpOnly: true,
				SameSite: http.SameSiteStrictMode,
				MaxAge:   -1,
			})
		case core.EFORBIDDEN:
			w.WriteHeader(http.StatusForbidden)
		}

		writeErr := writeJSON(w, Error{
			TraceId:  r.Context().Value("traceId").(string),
			Success:  false,
			Messages: []string{i18nError.ErrMsg()},
		})
		if writeErr != nil {
			l.Error("failed to write json", zap.Error(writeErr))
		}
		if i18nError.Code == core.EINTERNAL {
			logError(l, r, err, http.StatusInternalServerError)
		}
		return
	}

	i18nErrors, ok := err.(*core.I18NErrors)
	if ok {
		var messages []string
		for _, i18nError = range i18nErrors.Errors {
			switch i18nError.Code {
			case core.EINTERNAL:
				w.WriteHeader(http.StatusInternalServerError)
			case core.ENOTFOUND:
				w.WriteHeader(http.StatusNotFound)
			case core.EINVALID:
				w.WriteHeader(http.StatusBadRequest)
			case core.ECONFLICT:
				w.WriteHeader(http.StatusConflict)
			case core.EUNAUTHORIZED:
				w.WriteHeader(http.StatusUnauthorized)
				http.SetCookie(w, &http.Cookie{
					Name:     "ardo-app-session",
					Value:    "",
					Path:     "/",
					HttpOnly: true,
					SameSite: http.SameSiteStrictMode,
					MaxAge:   -1,
				})
			case core.EFORBIDDEN:
				w.WriteHeader(http.StatusForbidden)
			}
			messages = append(messages, i18nError.ErrMsg())
			if i18nError.Code == core.EINTERNAL {
				logError(l, r, err, http.StatusInternalServerError)
			}
		}
		writeErr := writeJSON(w, Error{
			TraceId:  r.Context().Value("traceId").(string),
			Success:  false,
			Messages: messages,
		})
		if writeErr != nil {
			l.Error("failed to write json", zap.Error(writeErr))
		}
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	writeErr := writeJSON(w, Error{
		TraceId:  r.Context().Value("traceId").(string),
		Success:  false,
		Messages: []string{err.Error()},
	})
	if writeErr != nil {
		l.Error("failed to write json", zap.Error(writeErr))
	}
	logError(l, r, err, http.StatusInternalServerError)
}
