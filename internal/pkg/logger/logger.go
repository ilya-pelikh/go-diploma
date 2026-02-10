package logger

import (
	"diploma/internal/pkg/config"

	"go.uber.org/zap"
)

var Logger *zap.Logger = zap.NewNop()

func Load(env config.ENV) error {
	var err error

	switch env {
	case config.PRODUCTION:
		Logger, err = zap.NewProduction()
	case config.DEVELOPMENT:
		Logger, err = zap.NewDevelopment()
	default:
		Logger = zap.NewNop()
	}

	return err
}

func Get() *zap.Logger {
	return Logger
}
