package app

import (
	"context"

	"diploma/internal/pkg/api"
	"diploma/internal/pkg/db"
	"diploma/internal/pkg/env"
	"diploma/internal/pkg/logger"
)

func Run(ctx context.Context) {
	logger.Load()

	db.Init(env.TODO_DBFILE)

	server := api.Create()

	go func() {
		<-ctx.Done()
		server.Shutdown()
	}()

	server.Start()
}
