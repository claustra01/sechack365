package main

import (
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/handler"
)

func setupRouter(r *framework.Router) error {
	// TODO: check error
	api := r.Group("/api/v1")

	wk := api.Group("/.well-known")
	wk.Get("/nodeinfo", handler.Nodeinfo)
	wk.Get("/webfinger", handler.Webfinger)

	// mock actor endpoint
	api.Get("/actor/mock", handler.MockActor)

	// api.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello, World!")
	// })

	// api.Post("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Goodbye, World!")
	// })

	// api.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	id := r.PathValue("id")
	// 	fmt.Fprintf(w, "Hello, %s!", id)
	// })

	// hello := api.Group("/hello")
	// hello.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello!")
	// })

	return nil
}
