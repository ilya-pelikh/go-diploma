package api

import (
	"context"
	"net/http"
	"time"

	"diploma/internal/pkg/env"
)

type Server struct {
	server *http.Server
}

func registerRoutes(mux *http.ServeMux) {
	mux.Handle("/", http.FileServer(http.Dir("../web")))
	mux.Handle("/api/task", http.HandlerFunc(handleTask))
	mux.Handle("/api/tasks", http.HandlerFunc(handleTasks))
}

func Create() *Server {
	mux := http.NewServeMux()
	registerRoutes(mux)

	return &Server{
		server: &http.Server{
			Addr:         ":" + env.TODO_PORT,
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
