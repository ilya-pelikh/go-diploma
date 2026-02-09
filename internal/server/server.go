package http

import (
	"context"
	"diploma/internal/config"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func New() *Server {

	mux := http.NewServeMux()
	registerRoutes(mux)

	fmt.Println(config.Get().TODO_PORT)

	return &Server{
		server: &http.Server{
			Addr:         ":" + config.Get().TODO_PORT,
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
