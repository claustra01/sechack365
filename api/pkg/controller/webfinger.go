package controller

import (
	"github.com/claustra01/sechack365/pkg/openapi"
	"github.com/claustra01/sechack365/pkg/usecase"
)

type WebfingerController struct {
	WebfingerUsecase usecase.WebfingerUsecase
}

func NewWebfingerController() *WebfingerController {
	return &WebfingerController{
		WebfingerUsecase: usecase.WebfingerUsecase{},
	}
}

func (c *WebfingerController) NewWebfingerActorLinks(host, id, name string) *openapi.WellknownWebfinger {
	return c.WebfingerUsecase.NewWebfingerActorLinks(host, id, name)
}

func (c *WebfingerController) NewNodeInfoLinks(host string) *openapi.WellknownNodeinfo {
	return c.WebfingerUsecase.NewNodeInfoLinks(host)
}
