package web

import (
	"net/http"
	"strings"
	"time"
)

type RequestInfo struct {
	Timestamp    time.Time
	RemoteAddr   string
	Referrer     string
	UserAgentRaw string
	UserAgent    UserAgent
}

func GetRequestInfo(r *http.Request) RequestInfo {
	headers := map[string][]string(r.Header)
	userAgentH, found := headers["User-Agent"]
	var userAgentS string
	if found {
		userAgentS = strings.Join(userAgentH, ", ")
	}
	return RequestInfo{
		Timestamp:    time.Now(),
		RemoteAddr:   r.RemoteAddr,
		UserAgentRaw: userAgentS,
		UserAgent:    ParseUA(userAgentS),
	}
}
