package logger

import (
	"diploma/internal/pkg/env"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger = zap.NewNop()

func Load() error {
	var err error

	switch env.MODE {
	case env.PRODUCTION:
		Logger, err = zap.NewProduction()
	case env.DEVELOPMENT:
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
