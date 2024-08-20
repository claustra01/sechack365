package usecase

import "github.com/claustra01/sechack365/pkg/model"

type IUserRepository interface {
	FindAll() ([]*model.User, error)
	FindById(id string) (*model.User, error)
	FindByUserId(string) (*model.User, error)
}

type UserUsecase struct {
	UserRepository IUserRepository
}

func (u *UserUsecase) FindAll() ([]*model.User, error) {
	return u.UserRepository.FindAll()
}

func (u *UserUsecase) FindById(id string) (*model.User, error) {
	return u.UserRepository.FindById(id)
}

func (u *UserUsecase) FindByUserId(userId string) (*model.User, error) {
	return u.UserRepository.FindByUserId(userId)
}
