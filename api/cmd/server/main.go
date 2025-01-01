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
	conn, err := infrastructure.NewSqlHandler(connStr)
	if err != nil {
		panic(err)
	}

	// Context
	ctx := framework.NewContext(logger, conn)

	// Websocket Connection
	relays, err := ctx.Controllers.NostrRelay.FindAll()
	if err != nil {
		panic(err)
	}
	ws, err := infrastructure.NewWsHandler(relays, logger)
	if err != nil {
		panic(err)
	}
	ctx.SetNostrRelays(ws)

	// Server
	server := framework.NewServer(ctx)
	router := server.Router

	router.Use(framework.LoggingMiddleware(logger), framework.RecoverMiddleware(logger), framework.CorsMiddleware(ctx.Config.Host))
	if err := setupRouter(router, logger); err != nil {
		panic(err)
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
