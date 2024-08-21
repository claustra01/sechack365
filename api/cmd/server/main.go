package main

import (
	"os"

	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/infrastructure"
)

func main() {
	// Logger
	logger := infrastructure.NewLogger(os.Getenv("LOG_LEVEL"))

	// DB Connection
	connStr := os.Getenv("POSTGRES_URL")
	// FIXME: なぜか?sslmode=disableだけ環境変数から読み取れない
	conn, err := infrastructure.NewSqlHandler(connStr + "?sslmode=disable")
	if err != nil {
		panic(err)
	}

	ctx := framework.NewContext(logger, conn)
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
