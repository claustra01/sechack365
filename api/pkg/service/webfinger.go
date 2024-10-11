package service

import (
	"fmt"

	"github.com/claustra01/sechack365/pkg/openapi"
)

type WebfingerService struct{}

func (s *WebfingerService) NewWebfingerActorLinks(host, id, name string) *openapi.WellknownWebfinger {
	var apService ActivitypubService
	return &openapi.WellknownWebfinger{
		Subject: fmt.Sprintf("acct:%s@%s", name, host),
		Links: []openapi.WellknownWebfingerLink{
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: apService.NewActorUrl(host, id),
			},
		},
	}
}

func (s *WebfingerService) NewNodeInfoLinks(host string) *openapi.WellknownNodeinfo {
	return &openapi.WellknownNodeinfo{
		Links: []openapi.WellknownNodeinfoLink{
			{
				Rel:  "http://nodeinfo.diaspora.software/ns/schema/2.0",
				Href: fmt.Sprintf("https://%s/api/v1/nodeinfo/2.0", host),
			},
		},
	}
}
