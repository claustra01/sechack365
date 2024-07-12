package main

import (
	"os"

	"github.com/claustra01/sechack365/pkg/framework"
)

func main() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "1323"
	}

	router := framework.NewRouter()
	if err := SetupRouter(router); err != nil {
		panic(err)
	}

	server := framework.NewServer(*router, port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
