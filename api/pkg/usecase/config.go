package usecase

import "github.com/claustra01/sechack365/pkg/model"

type INostrRelayHandler interface {
	FindAll() ([]*model.NostrRelay, error)
	Create(url string) (*model.NostrRelay, error)
	Delete(id string) error
}

type NostrRelayUsecase struct {
	NostrRelayHandler INostrRelayHandler
}

func (u *NostrRelayUsecase) FindAll() ([]*model.NostrRelay, error) {
	return u.NostrRelayHandler.FindAll()
}

func (u *NostrRelayUsecase) Create(url string) (*model.NostrRelay, error) {
	return u.NostrRelayHandler.Create(url)
}

func (u *NostrRelayUsecase) Delete(id string) error {
	return u.NostrRelayHandler.Delete(id)
}
