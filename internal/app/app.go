package app

import (
	"context"

	"diploma/internal/pkg/config"
	"diploma/internal/pkg/db"
	"diploma/internal/pkg/logger"
	httpserver "diploma/internal/server"
)

func Run(ctx context.Context) {
	env := config.Get()

	logger.Load(env.MODE)
	db.Init(env.TODO_DBFILE)

	server := httpserver.Create(env.TODO_PORT)

	go func() {
		<-ctx.Done()
		server.Shutdown()
	}()

	server.Start()
}
