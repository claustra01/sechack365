package usecase

import (
	"crypto/rsa"

	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/openapi"
)

// TODO: この型はmodelかopenapiに移動する
type FollowActivity struct {
	Context any    `json:"@context"`
	Type    string `json:"type"`
	Id      string `json:"id"`
	Actor   string `json:"actor"`
	Object  string `json:"object"`
}

type IActivityPubService interface {
	NewActor(user model.UserWithIdentifiers) *openapi.Actor
	NewActorUrl(host, id string) string
	NewKeyIdUrl(host, name string) string
	NewFollowActivity(id, host, followerId, followeeUrl string) *FollowActivity
	NewNodeInfo(userUsage int) *openapi.Nodeinfo
	ResolveWebfinger(username, host string) (string, error)
	ResolveRemoteActor(link string) (*openapi.Actor, error)
	SendActivity(keyId string, privKey *rsa.PrivateKey, targetHost string, activity any) ([]byte, error)
}

type ActivityPubUsecase struct {
	ActivityPubService IActivityPubService
}

func (u *ActivityPubUsecase) NewActor(user model.UserWithIdentifiers) *openapi.Actor {
	return u.ActivityPubService.NewActor(user)
}

func (u *ActivityPubUsecase) NewActorUrl(host, id string) string {
	return u.ActivityPubService.NewActorUrl(host, id)
}

func (u *ActivityPubUsecase) NewKeyIdUrl(host string, name string) string {
	return u.ActivityPubService.NewKeyIdUrl(host, name)
}

func (u *ActivityPubUsecase) NewFollowActivity(id, host, followerId, followeeUrl string) *FollowActivity {
	return u.ActivityPubService.NewFollowActivity(id, host, followerId, followeeUrl)
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

func (u *ActivityPubUsecase) SendActivity(keyId string, privKey *rsa.PrivateKey, targetHost string, activity any) ([]byte, error) {
	return u.ActivityPubService.SendActivity(keyId, privKey, targetHost, activity)
}
