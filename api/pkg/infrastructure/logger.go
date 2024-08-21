package infrastructure

import (
	"log/slog"
	"os"

	"github.com/claustra01/sechack365/pkg/model"
)

type Logger struct {
	Logger *slog.Logger
	Level  slog.Level
}

func NewLogger(levelStr string) model.ILogger {
	level := convertLogLevel(levelStr)
	slog.SetLogLoggerLevel(level)
	logger := new(Logger)
	logger.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Level = level
	return logger
}

func (l *Logger) Debug(msg string, args ...any) {
	l.Logger.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.Logger.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.Logger.Error(msg, args...)
}

func (l *Logger) Fatal(msg string, args ...any) {
	l.Logger.Error(msg, args...)
	os.Exit(1)
}

func convertLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
