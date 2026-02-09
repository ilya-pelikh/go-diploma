package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"diploma/internal/app"
	"diploma/internal/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	config.Load()

	app.Run(ctx)
}
