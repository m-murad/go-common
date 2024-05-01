package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
	appName string
}

func New(appName string, logLevel zapcore.Level) *Logger {
	zl := zap.Must(zap.NewProduction()).With(zap.String("app", appName))

	l := &Logger{
		Logger:  zl,
		appName: appName,
	}

	return l
}

func NewNoOp() *Logger {
	return &Logger{
		Logger: zap.NewNop(),
	}
}
