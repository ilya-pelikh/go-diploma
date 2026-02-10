package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"diploma/internal/app"
	"diploma/internal/pkg/env"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	env.Load()

	app.Run(ctx)
}
