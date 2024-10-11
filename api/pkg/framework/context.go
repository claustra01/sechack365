package framework

import (
	"context"

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
	Transaction      *controller.TransactionController
	User             *controller.UserController
	ApUserIdentifier *controller.ApUserIdentifierController
	Follow           *controller.FollowController
	ActivityPub      *controller.ActivityPubController
	Nostr            *controller.NostrController
	Webfinger        *controller.WebfingerController
}

func NewContext(logger model.ILogger, conn model.ISqlHandler, ws model.IWsHandler) *Context {
	return &Context{
		Ctx:         context.Background(),
		Logger:      logger,
		Config:      NewConfig(logger),
		Controllers: NewControllers(conn, ws),
	}
}

func NewControllers(conn model.ISqlHandler, ws model.IWsHandler) *Controllers {
	return &Controllers{
		Transaction:      controller.NewTransactionController(conn),
		User:             controller.NewUserController(conn),
		ApUserIdentifier: controller.NewApUserIdentifierController(conn),
		Follow:           controller.NewFollowController(conn),
		ActivityPub:      controller.NewActivityPubController(),
		Nostr:            controller.NewNostrController(ws),
		Webfinger:        controller.NewWebfingerController(),
	}
}
