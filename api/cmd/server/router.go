package main

import (
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/handler"
)

func setupRouter(r *framework.Router) error {
	api := r.Group("/api/v1")

	users := api.Group("/users")
	users.Get("", handler.GetAllUsers)
	users.Get("/{id}", handler.GetUser)
	users.Post("/{id}/inbox", handler.ActorInbox)
	users.Post("/{id}/outbox", handler.ActorOutbox)

	follow := api.Group("/follows")
	follow.Post("", handler.CreateFollow)

	api.Get("/lookup/{username}", handler.LookupUser)

	ni := api.Group("/nodeinfo")
	ni.Get("/2.0", handler.Nodeinfo2_0)

	wk := r.Group("/.well-known")
	wk.Get("/nodeinfo", handler.NodeinfoLinks)
	wk.Get("/webfinger", handler.WebfingerLinks)

	return nil
}
