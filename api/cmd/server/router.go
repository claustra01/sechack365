package main

import (
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/handler"
	"github.com/claustra01/sechack365/pkg/model"
)

func setupRouter(r *framework.Router, lg model.ILogger) error {
	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	if err := auth.Post("/register", handler.Register); err != nil {
		return err
	}
	if err := auth.Post("/login", handler.Login); err != nil {
		return err
	}
	if err := auth.Post("/logout", handler.Logout); err != nil {
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
	if err := users.Get("/{id}/follows", handler.GetUserFollows); err != nil {
		return err
	}
	if err := users.Get("/{id}/followers", handler.GetUserFollowers); err != nil {
		return err
	}
	if err := users.Get("/{id}/posts", handler.GetUserPosts); err != nil {
		return err
	}

	follow := api.Group("/follows")
	if err := follow.Post("", handler.CreateFollow); err != nil {
		return err
	}
	if err := follow.Delete("", handler.DeleteFollow); err != nil {
		return err
	}
	if err := follow.Get("/following/{id}", handler.CheckIsFollowing); err != nil {
		return err
	}

	if err := api.Get("/lookup/{username}", handler.LookupUser); err != nil {
		return err
	}

	post := api.Group("/posts")
	if err := post.Post("", handler.CreatePost, framework.AuthMiddleware(lg)); err != nil {
		return err
	}
	if err := post.Get("/{id}", handler.GetPost); err != nil {
		return err
	}

	article := api.Group("/articles")
	if err := article.Post("", handler.CreateArticle, framework.AuthMiddleware(lg)); err != nil {
		return err
	}
	if err := article.Get("/{id}", handler.GetArticle); err != nil {
		return err
	}
	if err := article.Get("/{id}/comments", handler.GetArticleCommentsById); err != nil {
		return err
	}

	if err := api.Get("/timeline", handler.GetTimeline); err != nil {
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
	if err := wk.Get("/nostr.json", handler.Nip05); err != nil {
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
