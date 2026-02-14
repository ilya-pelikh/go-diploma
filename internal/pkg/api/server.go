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
	mux.Handle("/", http.FileServer(http.Dir("web")))
	mux.Handle("/api/signin", http.HandlerFunc(handleAuth))
	mux.Handle("/api/task", auth(http.HandlerFunc(handleTask)))
	mux.Handle("/api/task/done", auth(http.HandlerFunc(handleTaskDone)))
	mux.Handle("/api/tasks", auth(http.HandlerFunc(handleTasks)))
	mux.Handle("/api/nextdate", http.HandlerFunc(handlePlanner))
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
