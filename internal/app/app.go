package app

import (
	"context"

	httpserver "diploma/internal/server"
)

func Run(ctx context.Context) {
	server := httpserver.New()

	go func() {
		<-ctx.Done()
		server.Shutdown()
	}()

	server.Start()
}
