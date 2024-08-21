package framework

import (
	"os"

	"github.com/claustra01/sechack365/pkg/model"
)

type Config struct {
	Host string
	Port string
}

func NewConfig(logger model.ILogger) *Config {
	var host string
	if host = os.Getenv("HOST"); host == "" {
		logger.Warn("HOST is not set. Using default value.")
		host = "localhost"
	}

	var port string
	if port = os.Getenv("PORT"); port == "" {
		logger.Warn("PORT is not set. Using default value.")
		port = "1323"
	}

	return &Config{
		Host: host,
		Port: port,
	}
}
