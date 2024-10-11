package usecase

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/openapi"
)

type IActivityPubService interface {
	NewActor(user model.User, identifier model.ApUserIdentifier) *openapi.Actor
	NewActorUrl(host, id string) string
	NewKeyIdUrl(host string, name string) string
	NewNodeInfo(userUsage int) *openapi.Nodeinfo
	ResolveWebfinger(username, host string) (string, error)
	ResolveRemoteActor(link string) (*openapi.Actor, error)
}

type ActivityPubUsecase struct {
	ActivityPubService IActivityPubService
}

func (u *ActivityPubUsecase) NewActor(user model.User, identifier model.ApUserIdentifier) *openapi.Actor {
	return u.ActivityPubService.NewActor(user, identifier)
}

func (u *ActivityPubUsecase) NewActorUrl(host, id string) string {
	return u.ActivityPubService.NewActorUrl(host, id)
}

func (u *ActivityPubUsecase) NewKeyIdUrl(host string, name string) string {
	return u.ActivityPubService.NewKeyIdUrl(host, name)
}

func (u *ActivityPubUsecase) NewNodeInfo(userUsage int) *openapi.Nodeinfo {
	return u.ActivityPubService.NewNodeInfo(userUsage)
}

func (u *ActivityPubUsecase) ResolveWebfinger(username, host string) (string, error) {
	return u.ActivityPubService.ResolveWebfinger(username, host)
}

func (u *ActivityPubUsecase) ResolveRemoteActor(link string) (*openapi.Actor, error) {
	return u.ActivityPubService.ResolveRemoteActor(link)
}
