package log

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var loggerOnce sync.Once

type Field = zapcore.Field

func New() {
	loggerOnce.Do(func() {
		Logger, _ = zap.NewProduction()
	})
}

func Sync() {
	Logger.Sync()
}

func Log() *zap.Logger {
	return Logger
}

func Error(msg string, fields ...Field) {
	Logger.Error(msg, fields...)
}

func Fatal(msg string, err error) {
	Logger.Fatal(msg, zap.Error(err))
}

func Info(msg string, fields ...Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	Logger.Warn(msg, fields...)
}
