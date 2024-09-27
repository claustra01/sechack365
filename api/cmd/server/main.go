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

	// Nostr Relay Connection
	// TODO: 接続するRelayを自由に設定できるようにする
	ws, err := infrastructure.NewWsHandler([]string{"wss://yabu.me"})
	if err != nil {
		panic(err)
	}
	defer ws.Close()

	ctx := framework.NewContext(logger, conn, ws)
	server := framework.NewServer(ctx)
	router := server.Router

	router.Use(framework.LoggingMiddleware(logger), framework.RecoverMiddleware(logger))
	if err := setupRouter(router); err != nil {
		panic(err)
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
