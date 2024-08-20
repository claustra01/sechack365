package main

import (
	"os"

	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/infrastructure"
)

func main() {
	connStr := os.Getenv("POSTGRES_URL")
	conn, err := infrastructure.NewSqlHandler(connStr)
	if err != nil {
		panic(err)
	}

	ctx := framework.NewContext(conn)
	server := framework.NewServer(ctx)
	router := server.Router

	router.Use(framework.LoggingMiddleware, framework.RecoverMiddleware)
	if err := setupRouter(router); err != nil {
		panic(err)
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
