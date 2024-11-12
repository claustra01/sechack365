package main

import (
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/handler"
	"github.com/claustra01/sechack365/pkg/model"
)

func setupRouter(r *framework.Router, lg model.ILogger) error {
	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	if err := auth.Post("/login", handler.Login); err != nil {
		return err
	}
	if err := auth.Post("/logout", handler.Logout, framework.AuthMiddleware(lg)); err != nil {
		return err
	}

	users := api.Group("/users")
	if err := users.Get("", handler.GetAllUsers); err != nil {
		return err
	}
	if err := users.Get("/me", handler.GetCurrentUser, framework.AuthMiddleware(lg)); err != nil {
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

	dev := api.Group("/dev", framework.DevApiMiddleware(lg))
	if err := dev.Get("/mock", handler.GenerateMock); err != nil {
		return err
	}
	if err := dev.Get("/reset", handler.ResetMock); err != nil {
		return err
	}

	return nil
}
