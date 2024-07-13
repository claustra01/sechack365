package framework

import (
	"context"
	"log/slog"
	"os"
)

type Context struct {
	Ctx    context.Context
	Config *Config
}

type Config struct {
	Host     string
	Port     string
	LogLevel slog.Level
}

func NewContext() *Context {
	return &Context{
		Ctx:    context.Background(),
		Config: NewConfig(),
	}
}

func NewConfig() *Config {
	var host string
	if host = os.Getenv("HOST"); host == "" {
		slog.Warn("HOST is not set. Using default value.")
		host = "localhost"
	}

	var port string
	if port = os.Getenv("PORT"); port == "" {
		slog.Warn("PORT is not set. Using default value.")
		port = "1323"
	}

	var logLevel string
	if logLevel = os.Getenv("LOG_LEVEL"); logLevel == "" {
		slog.Warn("LOG_LEVEL is not set. Using default value.")
		logLevel = "info"
	}

	return &Config{
		Host:     host,
		Port:     port,
		LogLevel: convertLogLevel(logLevel),
	}
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
