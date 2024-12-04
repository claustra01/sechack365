package usecase

import "github.com/claustra01/sechack365/pkg/model"

type IUserRepository interface {
	CreateLocalUser(username, password, displayName, profile, icon, host string) (*model.UserWithIdentifiers, error)
	CreateRemoteApUser(user *model.User, identifier *model.ApUserIdentifier) (*model.UserWithIdentifiers, error)
	CreateRemoteNostrUser(user *model.User, identifier *model.NostrUserIdentifier) (*model.UserWithIdentifiers, error)
	FindAll() ([]*model.UserWithIdentifiers, error)
	FindById(id string) (*model.UserWithIdentifiers, error)
	FindByLocalUsername(username string) (*model.UserWithIdentifiers, error)
	FindByApUsername(username string, host string) (*model.UserWithIdentifiers, error)
	FindByNostrPublicKey(publicKey string) (*model.UserWithIdentifiers, error)
	UpdateRemoteApUser(user *model.User, identifier *model.ApUserIdentifier) (*model.UserWithIdentifiers, error)
	UpdateRemoteNostrUser(user *model.User, identifier *model.NostrUserIdentifier) (*model.UserWithIdentifiers, error)
	DeleteById(id string) error
}

type UserUsecase struct {
	UserRepository IUserRepository
}

func (u *UserUsecase) CreateLocalUser(username, password, displayName, profile, icon, host string) (*model.UserWithIdentifiers, error) {
	return u.UserRepository.CreateLocalUser(username, password, displayName, profile, icon, host)
}

func (u *UserUsecase) CreateRemoteApUser(user *model.User, identifier *model.ApUserIdentifier) (*model.UserWithIdentifiers, error) {
	return u.UserRepository.CreateRemoteApUser(user, identifier)
}

func (u *UserUsecase) CreateRemoteNostrUser(user *model.User, identifier *model.NostrUserIdentifier) (*model.UserWithIdentifiers, error) {
	return u.UserRepository.CreateRemoteNostrUser(user, identifier)
}

func (u *UserUsecase) FindAll() ([]*model.UserWithIdentifiers, error) {
	return u.UserRepository.FindAll()
}

func (u *UserUsecase) FindById(id string) (*model.UserWithIdentifiers, error) {
	return u.UserRepository.FindById(id)
}

func (u *UserUsecase) FindByLocalUsername(username string) (*model.UserWithIdentifiers, error) {
	return u.UserRepository.FindByLocalUsername(username)
}

func (u *UserUsecase) FindByApUsername(username string, host string) (*model.UserWithIdentifiers, error) {
	return u.UserRepository.FindByApUsername(username, host)
}

func (u *UserUsecase) FindByNostrPublicKey(publicKey string) (*model.UserWithIdentifiers, error) {
	return u.UserRepository.FindByNostrPublicKey(publicKey)
}

func (u *UserUsecase) UpdateRemoteApUser(user *model.User, identifier *model.ApUserIdentifier) (*model.UserWithIdentifiers, error) {
	return u.UserRepository.UpdateRemoteApUser(user, identifier)
}

func (u *UserUsecase) UpdateRemoteNostrUser(user *model.User, identifier *model.NostrUserIdentifier) (*model.UserWithIdentifiers, error) {
	return u.UserRepository.UpdateRemoteNostrUser(user, identifier)
}

func (u *UserUsecase) DeleteById(id string) error {
	return u.UserRepository.DeleteById(id)
}
