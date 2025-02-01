package main

import (
	"os"

	"github.com/claustra01/sechack365/cmd/batch"
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/infrastructure"
)

func main() {
	// Logger
	logger := infrastructure.NewLogger(os.Getenv("LOG_LEVEL"))

	// DB Connection
	dbConnStr := os.Getenv("POSTGRES_URL")
	dbConn, err := infrastructure.NewSqlHandler(dbConnStr)
	if err != nil {
		panic(err)
	}

	// Storage Connection
	minioHost := os.Getenv("MINIO_HOST")
	minioPort := os.Getenv("MINIO_PORT")
	minioUser := os.Getenv("MINIO_ROOT_USER")
	minioPassword := os.Getenv("MINIO_ROOT_PASSWORD")
	storageConn, err := infrastructure.NewStorageHandler(minioHost, minioPort, minioUser, minioPassword)
	if err != nil {
		panic(err)
	}

	// Context
	ctx := framework.NewContext(logger, dbConn, storageConn)

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
	defer ws.Close()

	// Batch
	batch.UpdateNostrRemotePosts(ctx)

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
