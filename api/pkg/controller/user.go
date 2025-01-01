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

func (c *UserController) CreateLocalUser(username, password, displayName, profile, icon, host string) error {
	return c.UserUsecase.CreateLocalUser(username, password, displayName, profile, icon, host)
}

func (c *UserController) CreateRemoteApUser(user *model.User, identifier *model.ApUserIdentifier) error {
	return c.UserUsecase.CreateRemoteApUser(user, identifier)
}

func (c *UserController) CreateRemoteNostrUser(user *model.User, identifier *model.NostrUserIdentifier) error {
	return c.UserUsecase.CreateRemoteNostrUser(user, identifier)
}

func (c *UserController) FindAll() ([]*model.UserWithIdentifiers, error) {
	return c.UserUsecase.FindAll()
}

func (c *UserController) FindById(id string) (*model.UserWithIdentifiers, error) {
	return c.UserUsecase.FindById(id)
}

func (c *UserController) FindByLocalUsername(username string) (*model.UserWithIdentifiers, error) {
	return c.UserUsecase.FindByLocalUsername(username)
}

func (c *UserController) FindByApUsername(username string, host string) (*model.UserWithIdentifiers, error) {
	return c.UserUsecase.FindByApUsername(username, host)
}

func (c *UserController) FindByNostrNpub(publicKey string) (*model.UserWithIdentifiers, error) {
	return c.UserUsecase.FindByNostrNpub(publicKey)
}

func (c *UserController) UpdateRemoteApUser(user *model.User, identifier *model.ApUserIdentifier) error {
	return c.UserUsecase.UpdateRemoteApUser(user, identifier)
}

func (c *UserController) UpdateRemoteNostrUser(user *model.User, identifier *model.NostrUserIdentifier) error {
	return c.UserUsecase.UpdateRemoteNostrUser(user, identifier)
}

func (c *UserController) DeleteById(id string) error {
	return c.UserUsecase.DeleteById(id)
}

func (c *UserController) FindWithHashedPassword(username string) (*model.User, error) {
	return c.UserUsecase.FindWithHashedPassword(username)
}

func (c *UserController) GetAllFollowingNostrPubKeys() ([]string, error) {
	return c.UserUsecase.GetAllFollowingNostrPubKeys()
}

func (c *UserController) GetNostrPrivKey(id string) (string, error) {
	return c.UserUsecase.GetNostrPrivKey(id)
}
