package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/claustra01/sechack365/pkg/framework"
)

type nodeinfo struct {
	OpenRegistrations bool     `json:"openRegistrations"`
	Protocols         []string `json:"protocols"`
	Software          software `json:"software"`
	Usage             usage    `json:"usage"`
	Services          services `json:"services"`
	Metadata          metadata `json:"metadata"`
	Version           string   `json:"version"`
}

type software struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type usage struct {
	Users usersUsage `json:"users"`
}

type usersUsage struct {
	Total int `json:"total"`
}

type services struct {
	Inbound  struct{} `json:"inbound"`
	Outbound struct{} `json:"outbound"`
}

type metadata struct{}

func Nodeinfo(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// mock nodeinfo
		nodeinfo := nodeinfo{
			OpenRegistrations: false,
			Protocols: []string{
				"activitypub",
			},
			Software: software{
				Name:    "sechack365",
				Version: "0.1.0",
			},
			Usage: usage{
				Users: usersUsage{
					Total: 1,
				},
			},
			Services: services{
				Inbound:  struct{}{},
				Outbound: struct{}{},
			},
			Metadata: metadata{},
			Version:  "2.1",
		}

		w.Header().Set("Content-Type", "application/json")
		data, err := json.Marshal(nodeinfo)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			panic(err)
		}
		fmt.Fprint(w, string(data))
	}
}

type webfinger struct {
	Subject string `json:"subject"`
	Links   []link `json:"links"`
}

type link struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	Href string `json:"href"`
}

func Webfinger(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resource := r.URL.Query().Get("resource")

		// mock actor
		exceptedResource := fmt.Sprintf("acct:%s@%s", "mock", c.Config.Host)
		if resource != exceptedResource {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		webfinger := webfinger{
			Subject: exceptedResource,
			Links: []link{
				{
					Rel:  "self",
					Type: "application/activity+json",
					Href: fmt.Sprintf("http://%s/api/v1/actor/mock", c.Config.Host),
				},
			},
		}

		data, err := json.Marshal(webfinger)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(data))
	}
}
