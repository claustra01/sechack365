package activitypub

import "fmt"

type Webfinger struct {
	Subject string `json:"subject"`
	Links   []Link `json:"links"`
}

type Link struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	Href string `json:"href"`
}

func GetWebfingerActor(name string, host string) *Webfinger {
	return &Webfinger{
		Subject: fmt.Sprintf("acct:%s@%s", name, host),
		Links: []Link{
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: fmt.Sprintf("https://%s/api/v1/actor/%s", host, name),
			},
		},
	}
}
