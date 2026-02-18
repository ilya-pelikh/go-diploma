package logger

import (
	"diploma/internal/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger = zap.NewNop()

func Load() error {
	var err error

	switch config.ENV.TODO_ENV {
	case "production":
		Logger, err = zap.NewProduction()
	case "development":
		Logger, err = zap.NewDevelopment(
			zap.AddStacktrace(zapcore.ErrorLevel),
		)
	default:
		Logger = zap.NewNop()
	}

	return err
}

func Get() *zap.Logger {
	return Logger
}
