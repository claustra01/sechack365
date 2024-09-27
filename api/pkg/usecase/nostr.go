package usecase

import "github.com/claustra01/sechack365/pkg/model"

type INostrService interface {
	GetUserProfile(id string) (*model.NostrProfile, error)
}

type NostrUsecase struct {
	NostrService INostrService
}

func (u *NostrUsecase) GetUserProfile(id string) (*model.NostrProfile, error) {
	return u.NostrService.GetUserProfile(id)
}
