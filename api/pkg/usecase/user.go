package usecase

import "github.com/claustra01/sechack365/pkg/model"

type IUserRepository interface {
	FindAll() ([]*model.User, error)
	FindById(id string) (*model.User, error)
	FindByUsername(username string, host string) (*model.User, error)
	Insert(username string, password string, host string, display_name string, profile string) (*model.User, error)
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

func (u *UserUsecase) FindByUsername(username string, host string) (*model.User, error) {
	return u.UserRepository.FindByUsername(username, host)
}

func (u *UserUsecase) Insert(username string, password string, host string, display_name string, profile string) (*model.User, error) {
	return u.UserRepository.Insert(username, password, host, display_name, profile)
}

type IApUserIdentifierRepository interface {
	Insert(userId string) (*model.ApUserIdentifier, error)
}

type ApUserIdentifierUsecase struct {
	ApUserIdentifierRepository IApUserIdentifierRepository
}

func (u *ApUserIdentifierUsecase) Insert(userId string) (*model.ApUserIdentifier, error) {
	return u.ApUserIdentifierRepository.Insert(userId)
}

type IApUserRepository interface {
	FindByUsername(username string, host string) (*model.ApUser, error)
}

type ApUserUsecase struct {
	ApUserRepository IApUserRepository
}

func (u *ApUserUsecase) FindByUsername(username string, host string) (*model.ApUser, error) {
	return u.ApUserRepository.FindByUsername(username, host)
}
