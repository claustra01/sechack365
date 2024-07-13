package main

import (
	"github.com/claustra01/sechack365/pkg/framework"
)

func main() {
	config := framework.NewServerConfig()
	server := framework.NewServer(config)
	router := server.Router

	router.Use(framework.LoggingMiddleware, framework.RecoverMiddleware)
	if err := setupRouter(router); err != nil {
		panic(err)
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
