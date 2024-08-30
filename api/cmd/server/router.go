package main

import (
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/handler"
)

func setupRouter(r *framework.Router) error {
	api := r.Group("/api/v1")

	users := api.Group("/users")
	users.Get("", handler.GetAllUsers)
	users.Get("/{username}", handler.LookupUser)

	actor := api.Group("/actor/{username}")
	actor.Get("", handler.GetActor)
	actor.Post("/inbox", handler.ActorInbox)
	actor.Post("/outbox", handler.ActorOutbox)

	ni := api.Group("/nodeinfo")
	ni.Get("/2.0", handler.Nodeinfo2_0)

	wk := r.Group("/.well-known")
	wk.Get("/nodeinfo", handler.NodeinfoLinks)
	wk.Get("/webfinger", handler.WebfingerLinks)

	return nil
}
