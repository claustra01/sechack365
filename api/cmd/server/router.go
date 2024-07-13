package main

import (
	"fmt"
	"net/http"

	"github.com/claustra01/sechack365/pkg/framework"
)

func SetupRouter(r *framework.Router) error {
	// TODO: check error
	api := r.Group("/api/v1")

	api.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	api.Post("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Goodbye, World!")
	})

	api.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Hello, %s!", id)
	})

	hello := api.Group("/hello")
	hello.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!")
	})

	r.HandleRoutes()
	return nil
}
