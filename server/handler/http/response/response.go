package response

import (
	"encoding/json"
	"github.com/nnniyaz/blog/pkg/logger"
	"github.com/nnniyaz/blog/pkg/web"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func writeJSON(w http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}

func logError(l logger.Logger, r *http.Request, err error, statusCode int) {
	traceId := r.Context().Value("traceId").(string)
	requestInfo := r.Context().Value("requestInfo").(web.RequestInfo)

	fields := []zap.Field{
		zap.String("traceId", traceId),
		zap.String("method", r.Method),
		zap.String("url", r.URL.String()),
		zap.String("proto", r.Proto),
		zap.String("remoteAddr", requestInfo.RemoteAddr),
		zap.String("host", r.Host),
		zap.String("statusCode", strconv.Itoa(statusCode)),
	}

	l.Error(err.Error(), fields...)
}
