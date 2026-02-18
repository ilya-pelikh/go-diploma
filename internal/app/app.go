package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"diploma/internal/pkg/config"
	"diploma/internal/pkg/db"
	"diploma/internal/pkg/logger"
	"diploma/internal/pkg/server"

	"go.uber.org/zap"
)

func Run(ctx context.Context) {
	err := logger.Load()
	if err != nil {
		fmt.Println("couldn't install logger")
	}

	err = db.Init(config.ENV.TODO_DBFILE)
	if err != nil {
		logger.Logger.Panic("couldn't install environment")
	}

	srv := server.Create()
	serveErr := make(chan error, 1)

	go func() {
		logger.Logger.Info("server is online")

		err := srv.Start()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			serveErr <- err
		}
	}()

	select {
	case <-ctx.Done():
	case err = <-serveErr:
		logger.Logger.Error("server start failed", zap.Error(err))
	}

	err = srv.Shutdown()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Logger.Error("server shutdown failed", zap.Error(err))
	}

	err = db.Close()
	if err != nil {
		logger.Logger.Error("database close failed", zap.Error(err))
	}

	logger.Logger.Info("server is offline", zap.Error(err))
}
