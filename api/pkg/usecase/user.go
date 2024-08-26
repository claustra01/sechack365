package usecase

import "github.com/claustra01/sechack365/pkg/model"

type IUserRepository interface {
	FindAll() ([]*model.User, error)
	FindById(id string) (*model.User, error)
	FindByUsername(string) (*model.User, error)
	Insert(username string, password string, host string) error
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

func (u *UserUsecase) FindByUsername(username string) (*model.User, error) {
	return u.UserRepository.FindByUsername(username)
}

func (u *UserUsecase) Insert(username string, password string, host string) error {
	return u.UserRepository.Insert(username, password, host)
}

type IApUserIdentifierRepository interface {
	Insert(userId string) error
}

type ApUserIdentifierUsecase struct {
	ApUserIdentifierRepository IApUserIdentifierRepository
}

type IApUserRepository interface {
	FindByUsername(string) (*model.ApUser, error)
}

type ApUserUsecase struct {
	ApUserRepository IApUserRepository
}

func (u *ApUserUsecase) FindByUsername(username string) (*model.ApUser, error) {
	return u.ApUserRepository.FindByUsername(username)
}
