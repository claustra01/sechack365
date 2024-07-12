package main

import (
	"fmt"
	"net/http"

	"github.com/claustra01/sechack365/pkg/framework"
)

func SetupRouter(r *framework.Router) error {
	// TODO: check error
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Goodbye, World!")
	})

	r.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Hello, %s!", id)
	})

	hello := r.Group("/hello")
	hello.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!")
	})

	r.HandleRoutes()
	return nil
}
