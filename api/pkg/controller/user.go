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

func (controller *UserController) FindAll() (users []*model.User, err error) {
	return controller.UserUsecase.FindAll()
}

func (controller *UserController) FindById(id string) (user *model.User, err error) {
	return controller.UserUsecase.FindById(id)
}

func (controller *UserController) FindByUsername(username string, host string) (user *model.User, err error) {
	return controller.UserUsecase.FindByUsername(username, host)
}

func (controller *UserController) Insert(username string, password string, host string, protocol string, display_name string, profile string, icon string) (*model.User, error) {
	return controller.UserUsecase.Insert(username, password, host, protocol, display_name, profile, icon)
}

func (controller *UserController) UpdateRemoteUser(username string, host string, display_name string, profile string, icon string) (*model.User, error) {
	return controller.UserUsecase.UpdateRemoteUser(username, host, display_name, profile, icon)
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

func (controller *ApUserIdentifierController) Insert(userId string, inbox string, outbox string, publicKey string) (*model.ApUserIdentifier, error) {
	return controller.ApUserIdentifierUsecase.Insert(userId, inbox, outbox, publicKey)
}

type ApUserController struct {
	ApUserUsecase usecase.ApUserUsecase
}

func NewApUserController(conn model.ISqlHandler) *ApUserController {
	return &ApUserController{
		ApUserUsecase: usecase.ApUserUsecase{
			ApUserRepository: &repository.ApUserRepository{
				SqlHandler: conn,
			},
		},
	}
}

func (controller *ApUserController) FindById(id string) (user *model.ApUser, err error) {
	return controller.ApUserUsecase.FindById(id)
}

func (controller *ApUserController) FindByUsername(username string, host string) (user *model.ApUser, err error) {
	return controller.ApUserUsecase.FindByUsername(username, host)
}
