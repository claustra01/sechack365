package main

import (
	"github.com/claustra01/sechack365/pkg/framework"
)

func main() {
	router := framework.NewRouter()
	router.Use(framework.LoggingMiddleware, framework.RecoverMiddleware)
	if err := setupRouter(router); err != nil {
		panic(err)
	}

	config := framework.NewServerConfig()
	server := framework.NewServer(*router, config)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
