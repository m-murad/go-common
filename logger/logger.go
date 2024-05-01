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
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(logLevel)
	cfg.EncoderConfig = encoderCfg

	zl := zap.Must(cfg.Build()).With(zap.String("app", appName))

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
