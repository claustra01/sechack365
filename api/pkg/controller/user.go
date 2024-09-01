package controller

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/repository"
	"github.com/claustra01/sechack365/pkg/usecase"
)

type UserController struct {
	UserUsecase usecase.UserUsecase
}

func NewUserController(conn model.ISqlHandler) *UserController {
	return &UserController{
		UserUsecase: usecase.UserUsecase{
			UserRepository: &repository.UserRepository{
				SqlHandler: conn,
			},
		},
	}
}

type ApUserIdentifierController struct {
	ApUserIdentifierUsecase usecase.ApUserIdentifierUsecase
}

func NewApUserIdentifierController(conn model.ISqlHandler) *ApUserIdentifierController {
	return &ApUserIdentifierController{
		ApUserIdentifierUsecase: usecase.ApUserIdentifierUsecase{
			ApUserIdentifierRepository: &repository.ApUserIdentifierRepository{
				SqlHandler: conn,
			},
		},
	}
}

func (c *UserController) FindAll() ([]*model.User, error) {
	return c.UserUsecase.FindAll()
}

func (c *UserController) FindById(id string) (*model.User, error) {
	return c.UserUsecase.FindById(id)
}

func (c *UserController) FindByUsername(username, host string) (*model.User, error) {
	return c.UserUsecase.FindByUsername(username, host)
}

func (c *UserController) CreateRemoteUser(username, host, protocol, displayName, profile, icon string) (*model.User, error) {
	return c.UserUsecase.CreateRemoteUser(username, host, protocol, displayName, profile, icon)
}

func (c *UserController) UpdateRemoteUser(username, host, displayName, profile, icon string) (*model.User, error) {
	return c.UserUsecase.UpdateRemoteUser(username, host, displayName, profile, icon)
}

func (c *ApUserIdentifierController) FindById(id string) (*model.ApUserIdentifier, error) {
	return c.ApUserIdentifierUsecase.FindById(id)
}
