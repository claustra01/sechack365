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
	User *controller.UserController
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
		User: controller.NewUserController(conn),
	}
}
