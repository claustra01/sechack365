package framework

import (
	"log/slog"
	"net/http"
	"os"
)

type ServerInterface interface {
	ListenAndServe() error
}

type Server struct {
	s      *http.Server
	Config *ServerConfig
	Router *Router
}

type ServerConfig struct {
	Host     string
	Port     string
	LogLevel slog.Level
}

func NewServer(config *ServerConfig) *Server {
	router := NewRouter()
	slog.SetLogLoggerLevel(config.LogLevel)

	return &Server{
		s: &http.Server{
			Addr:    ":" + config.Port,
			Handler: router.mux,
		},
		Config: config,
		Router: router,
	}
}

func (s *Server) ListenAndServe() error {
	return s.s.ListenAndServe()
}

func NewServerConfig() *ServerConfig {
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

	return &ServerConfig{
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
