package controller

import (
	"time"

	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/service"
	"github.com/claustra01/sechack365/pkg/usecase"
)

type NostrController struct {
	NostrUsecase usecase.NostrUsecase
}

func NewNostrController(ws model.IWsHandler) *NostrController {
	return &NostrController{
		NostrUsecase: usecase.NostrUsecase{
			NostrService: &service.NostrService{
				Ws: ws,
			},
		},
	}
}

func (c *NostrController) GetRemoteProfile(id string) (*model.NostrProfile, error) {
	return c.NostrUsecase.GetRemoteProfile(id)
}

func (c *NostrController) GetRemotePosts(pubKeys []string, since time.Time) ([]*model.NostrEvent, error) {
	return c.NostrUsecase.GetRemotePosts(pubKeys, since)
}

func (c *NostrController) PublishProfile(privKey string, profile *model.NostrProfile) error {
	return c.NostrUsecase.PublishProfile(privKey, profile)
}

func (c *NostrController) PublishPost(privKey string, note string) error {
	return c.NostrUsecase.PublishPost(privKey, note)
}

func (c *NostrController) PublishFollow(privKey string, pubKeys []string) error {
	return c.NostrUsecase.PublishFollow(privKey, pubKeys)
}
