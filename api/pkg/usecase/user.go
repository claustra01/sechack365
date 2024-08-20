package usecase

import "github.com/claustra01/sechack365/pkg/model"

type IUserRepository interface {
	FindAll() (users []*model.User, err error)
	FindById(id string) (user *model.User, err error)
	FindByUserId(userId string) (user *model.User, err error)
}

type UserUsecase struct {
	UserRepository IUserRepository
}

func (u *UserUsecase) FindAll() (users []*model.User, err error) {
	return u.UserRepository.FindAll()
}

func (u *UserUsecase) FindById(id string) (user *model.User, err error) {
	return u.UserRepository.FindById(id)
}

func (u *UserUsecase) FindByUserId(userId string) (user *model.User, err error) {
	return u.UserRepository.FindByUserId(userId)
}
