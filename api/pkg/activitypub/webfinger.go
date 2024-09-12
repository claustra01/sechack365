package activitypub

import "fmt"

type NodeInfoWebfinger struct {
	Links []NodeInfoLink `json:"links"`
}

type NodeInfoLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type Webfinger struct {
	Subject string          `json:"subject"`
	Links   []WebfingerLink `json:"links"`
}

type WebfingerLink struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	Href string `json:"href"`
}

func BuildWebfingerActorLinksSchema(host, id, name string) *Webfinger {
	return &Webfinger{
		Subject: fmt.Sprintf("acct:%s@%s", name, host),
		Links: []WebfingerLink{
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: BuildActorUrl(host, id),
			},
		},
	}
}

func BuildNodeInfoLinksSchema(host string) *NodeInfoWebfinger {
	return &NodeInfoWebfinger{
		Links: []NodeInfoLink{
			{
				Rel:  "http://nodeinfo.diaspora.software/ns/schema/2.0",
				Href: fmt.Sprintf("https://%s/api/v1/nodeinfo/2.0", host),
			},
		},
	}
}
