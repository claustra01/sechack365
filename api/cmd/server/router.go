package main

import (
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/handler"
)

func setupRouter(r *framework.Router) error {
	api := r.Group("/api/v1")

	users := api.Group("/users")
	if err := users.Get("", handler.GetAllUsers); err != nil {
		return err
	}
	if err := users.Get("/{id}", handler.GetUser); err != nil {
		return err
	}
	if err := users.Post("/{id}/inbox", handler.ActorInbox); err != nil {
		return err
	}
	if err := users.Post("/{id}/outbox", handler.ActorOutbox); err != nil {
		return err
	}

	follow := api.Group("/follows")
	if err := follow.Post("", handler.CreateFollow); err != nil {
		return err
	}

	if err := api.Get("/lookup/{username}", handler.LookupUser); err != nil {
		return err
	}

	ni := api.Group("/nodeinfo")
	if err := ni.Get("/2.0", handler.Nodeinfo2_0); err != nil {
		return err
	}

	wk := r.Group("/.well-known")
	if err := wk.Get("/nodeinfo", handler.NodeinfoLinks); err != nil {
		return err
	}
	if err := wk.Get("/webfinger", handler.WebfingerLinks); err != nil {
		return err
	}

	return nil
}
