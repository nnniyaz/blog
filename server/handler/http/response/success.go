package response

import (
	"go.uber.org/zap"
	"net/http"
	"scoring-system/pkg/logger"
)

type Success struct {
	TraceId string      `json:"traceId" example:"string"`
	Success bool        `json:"success" example:"true"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccess(l logger.Logger, w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeErr := writeJSON(w, Success{
		TraceId: r.Context().Value("traceId").(string),
		Success: true,
		Data:    data,
	})
	if writeErr != nil {
		l.Error("failed to write json", zap.Error(writeErr))
	}
}

func NewSuccessCreated(l logger.Logger, w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	writeErr := writeJSON(w, Success{
		TraceId: r.Context().Value("traceId").(string),
		Success: true,
		Data:    data,
	})
	if writeErr != nil {
		l.Error("failed to write json", zap.Error(writeErr))
	}
}
