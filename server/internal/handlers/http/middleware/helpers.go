package middleware

import (
	"net/http"
	"strconv"
)

func getParam(r *http.Request, name string) string {
	return r.URL.Query().Get(name)
}

func getParamInt64(r *http.Request, name string) (int64, error) {
	return strconv.ParseInt(getParam(r, name), 10, 64)
}

func getParamBool(r *http.Request, name string) (bool, error) {
	return strconv.ParseBool(getParam(r, name))
}
