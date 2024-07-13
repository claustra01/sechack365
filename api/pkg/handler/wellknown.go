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

	// Mock nodeinfo
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

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, err := json.Marshal(nodeinfo)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			panic(err)
		}
		fmt.Fprint(w, string(data))
	}
}

func Webfinger(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"subject":"acct:admin@localhost","links":[{"rel":"self","type":"application/activity+json","href":"http://localhost/api/v1/actor/admin"}]}`)
	}
}
