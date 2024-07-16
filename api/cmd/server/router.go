package main

import (
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/handler"
)

func setupRouter(r *framework.Router) error {
	wk := r.Group("/.well-known")
	wk.Get("/nodeinfo", handler.Nodeinfo)
	wk.Get("/webfinger", handler.Webfinger)

	api := r.Group("/api/v1")
	api.Get("/actor/{username}", handler.GetActor)

	return nil
}
