package controller

import (
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

func (c *NostrController) GetUserProfile(id string) (*model.NostrProfile, error) {
	return c.NostrUsecase.GetUserProfile(id)
}

func (c *NostrController) PostUserProfile(privKey string, pubKey string, profile *model.NostrProfile) error {
	return c.NostrUsecase.PostUserProfile(privKey, pubKey, profile)
}
