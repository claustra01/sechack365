package usecase

import (
	"time"

	"github.com/claustra01/sechack365/pkg/model"
)

type INostrService interface {
	GetRemoteProfile(id string) (*model.NostrProfile, error)
	GetRemotePosts(pubKeys []string, since time.Time) ([]*model.NostrEvent, error)
	PublishProfile(privKey string, profile *model.NostrProfile) error
	PublishPost(privKey string, note string) error
	PublishFollow(privKey string, pubKeys []string) error
}

type NostrUsecase struct {
	NostrService INostrService
}

func (u *NostrUsecase) GetRemoteProfile(id string) (*model.NostrProfile, error) {
	return u.NostrService.GetRemoteProfile(id)
}

func (u *NostrUsecase) GetRemotePosts(pubKeys []string, since time.Time) ([]*model.NostrEvent, error) {
	return u.NostrService.GetRemotePosts(pubKeys, since)
}

func (u *NostrUsecase) PublishProfile(privKey string, profile *model.NostrProfile) error {
	return u.NostrService.PublishProfile(privKey, profile)
}

func (u *NostrUsecase) PublishPost(privKey string, note string) error {
	return u.NostrService.PublishPost(privKey, note)
}

func (u *NostrUsecase) PublishFollow(privKey string, pubKeys []string) error {
	return u.NostrService.PublishFollow(privKey, pubKeys)
}
