package activitypub

import (
	"fmt"

	"github.com/claustra01/sechack365/pkg/openapi"
)

func BuildWebfingerActorLinksSchema(host, id, name string) *openapi.WellknownWebfinger {
	return &openapi.WellknownWebfinger{
		Subject: fmt.Sprintf("acct:%s@%s", name, host),
		Links: []openapi.WellknownWebfingerLink{
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: BuildActorUrl(host, id),
			},
		},
	}
}

func BuildNodeInfoLinksSchema(host string) *openapi.WellknownNodeinfo {
	return &openapi.WellknownNodeinfo{
		Links: []openapi.WellknownNodeinfoLink{
			{
				Rel:  "http://nodeinfo.diaspora.software/ns/schema/2.0",
				Href: fmt.Sprintf("https://%s/api/v1/nodeinfo/2.0", host),
			},
		},
	}
}
