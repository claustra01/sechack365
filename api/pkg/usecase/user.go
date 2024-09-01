package usecase

import "github.com/claustra01/sechack365/pkg/model"

type IUserRepository interface {
	FindAll() ([]*model.User, error)
	FindById(id string) (*model.User, error)
	FindByUsername(username string, host string) (*model.User, error)
	CreateRemoteUser(username, host, protocol, displayName, profile, icon string) (*model.User, error)
	UpdateRemoteUser(username, host, displayName, profile, icon string) (*model.User, error)
}

type UserUsecase struct {
	UserRepository IUserRepository
}

type IApUserIdentifierRepository interface {
	FindById(id string) (*model.ApUserIdentifier, error)
}

type ApUserIdentifierUsecase struct {
	ApUserIdentifierRepository IApUserIdentifierRepository
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

func (u *UserUsecase) CreateRemoteUser(username, host, protocol, displayName, profile, icon string) (*model.User, error) {
	return u.UserRepository.CreateRemoteUser(username, host, protocol, displayName, profile, icon)
}

func (u *UserUsecase) UpdateRemoteUser(username, host, displayName, profile, icon string) (*model.User, error) {
	return u.UserRepository.UpdateRemoteUser(username, host, displayName, profile, icon)
}

func (u *ApUserIdentifierUsecase) FindById(id string) (*model.ApUserIdentifier, error) {
	return u.ApUserIdentifierRepository.FindById(id)
}
