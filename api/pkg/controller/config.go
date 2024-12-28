package controller

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/repository"
	"github.com/claustra01/sechack365/pkg/usecase"
)

type NostrRelayController struct {
	NostrRelayUsecase usecase.NostrRelayUsecase
}

func NewNostrRelayController(conn model.ISqlHandler) *NostrRelayController {
	return &NostrRelayController{
		NostrRelayUsecase: usecase.NostrRelayUsecase{
			NostrRelayHandler: &repository.NostrRelayRepository{
				SqlHandler: conn,
			},
		},
	}
}

func (c *NostrRelayController) FindAll() ([]*model.NostrRelay, error) {
	return c.NostrRelayUsecase.FindAll()
}

func (c *NostrRelayController) Create(url string) error {
	return c.NostrRelayUsecase.Create(url)
}

func (c *NostrRelayController) Delete(id string) error {
	return c.NostrRelayUsecase.Delete(id)
}
