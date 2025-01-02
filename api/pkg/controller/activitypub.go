package controller

import (
	"crypto/rsa"

	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/openapi"
	"github.com/claustra01/sechack365/pkg/service"
	"github.com/claustra01/sechack365/pkg/usecase"
)

type ActivityPubController struct {
	ActivityPubUsecase usecase.ActivityPubUsecase
}

func NewActivityPubController() *ActivityPubController {
	return &ActivityPubController{
		ActivityPubUsecase: usecase.ActivityPubUsecase{
			ActivityPubService: &service.ActivitypubService{},
		},
	}
}

func (c *ActivityPubController) NewApContext() *openapi.Actor_Context {
	return c.ActivityPubUsecase.NewApContext()
}

func (c *ActivityPubController) NewActor(user model.UserWithIdentifiers) *openapi.Actor {
	return c.ActivityPubUsecase.NewActor(user)
}

func (c *ActivityPubController) NewActorUrl(host, id string) string {
	return c.ActivityPubUsecase.NewActorUrl(host, id)
}

func (c *ActivityPubController) NewKeyIdUrl(host string, name string) string {
	return c.ActivityPubUsecase.NewKeyIdUrl(host, name)
}

func (c *ActivityPubController) NewNodeInfo(userUsage int) *openapi.Nodeinfo {
	return c.ActivityPubUsecase.NewNodeInfo(userUsage)
}

func (c *ActivityPubController) ResolveWebfinger(username, host string) (string, error) {
	return c.ActivityPubUsecase.ResolveWebfinger(username, host)
}

func (c *ActivityPubController) ResolveRemoteActor(link string) (*openapi.Actor, error) {
	return c.ActivityPubUsecase.ResolveRemoteActor(link)
}

func (c *ActivityPubController) SendActivity(keyId string, privKey *rsa.PrivateKey, targetHost string, activity any) ([]byte, error) {
	return c.ActivityPubUsecase.SendActivity(keyId, privKey, targetHost, activity)
}
