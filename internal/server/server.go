package http

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func Create(port string) *Server {
	mux := http.NewServeMux()
	registerRoutes(mux)

	return &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      mux,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = s.server.Shutdown(ctx)
}
