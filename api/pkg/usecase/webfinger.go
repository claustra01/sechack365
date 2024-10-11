package usecase

import "github.com/claustra01/sechack365/pkg/openapi"

type IWebfingerUsecase interface {
	NewWebfingerActorLinks(host, id, name string) *openapi.WellknownWebfinger
	NewNodeInfoLinks(host string) *openapi.WellknownNodeinfo
}

type WebfingerUsecase struct {
	WebfingerService IWebfingerUsecase
}

func (u *WebfingerUsecase) NewWebfingerActorLinks(host, id, name string) *openapi.WellknownWebfinger {
	return u.WebfingerService.NewWebfingerActorLinks(host, id, name)
}

func (u *WebfingerUsecase) NewNodeInfoLinks(host string) *openapi.WellknownNodeinfo {
	return u.WebfingerService.NewNodeInfoLinks(host)
}
