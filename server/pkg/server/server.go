package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port int, handler http.Handler, startedAt time.Time) error {
	s.httpServer = &http.Server{
		Addr:           ":" + strconv.Itoa(port),
		Handler:        handler,
		MaxHeaderBytes: 5 << 20, // 5 MB
	}

	fmt.Println("+-----------------------------------------------+")
	fmt.Printf("| SERVER STARTED ON PORT: %d                  |\n", port)
	fmt.Printf("| Started at: %s         |\n", startedAt.Format("2006-01-02 15:04:05 -0700"))
	fmt.Println("+-----------------------------------------------+")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
