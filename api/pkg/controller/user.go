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

func (controller *UserController) FindByUsername(username string) (user *model.User, err error) {
	return controller.UserUsecase.FindByUsername(username)
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

func (controller *ApUserController) FindByUsername(username string) (user *model.ApUser, err error) {
	return controller.ApUserUsecase.FindByUsername(username)
}
