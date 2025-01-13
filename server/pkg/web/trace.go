package web

import (
	"github.com/google/uuid"
)

func GenerateTraceId() string {
	return uuid.New().String()
}
