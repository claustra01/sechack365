package framework

import (
	"context"
	"net/http"
	"time"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/controller"
	"github.com/claustra01/sechack365/pkg/model"
)

type Context struct {
	Ctx         context.Context
	Logger      model.ILogger
	Config      *Config
	Controllers *Controllers
}

type Controllers struct {
	Transaction *controller.TransactionController
	User        *controller.UserController
	Follow      *controller.FollowController
	Post        *controller.PostController
	Article     *controller.ArticleController
	NostrRelay  *controller.NostrRelayController
	ActivityPub *controller.ActivityPubController
	Nostr       *controller.NostrController
	Webfinger   *controller.WebfingerController
}

func NewContext(logger model.ILogger, conn model.ISqlHandler) *Context {
	return &Context{
		Ctx:         context.Background(),
		Logger:      logger,
		Config:      NewConfig(logger),
		Controllers: NewControllers(conn),
	}
}

func NewControllers(conn model.ISqlHandler) *Controllers {
	return &Controllers{
		Transaction: controller.NewTransactionController(conn),
		User:        controller.NewUserController(conn),
		Follow:      controller.NewFollowController(conn),
		Post:        controller.NewPostController(conn),
		NostrRelay:  controller.NewNostrRelayController(conn),
		ActivityPub: controller.NewActivityPubController(),
		// set websocket connection with SetNostrRelays()
		// Nostr:               controller.NewNostrController(ws),
		Webfinger: controller.NewWebfingerController(),
	}
}

func (c *Context) SetNostrRelays(ws model.IWsHandler) {
	c.Controllers.Nostr = controller.NewNostrController(ws)
}

func (c *Context) CurrentUser(r *http.Request) (*model.UserWithIdentifiers, error) {
	cookie, err := r.Cookie("session")
	if err != nil || cookie.Value == "" {
		return nil, cerror.ErrUserNotFound
	}
	sessionId := cookie.Value
	session, ok := Sessions[sessionId]
	if !ok || session.ExpiredAt.Before(time.Now()) {
		return nil, cerror.ErrUserNotFound
	}
	user, err := c.Controllers.User.FindById(session.UserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
