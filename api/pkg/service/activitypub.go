package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/openapi"
)

type ActivitypubService struct{}

// NOTE: const values: start

func NewApContext() openapi.Actor_Context {
	var ApContext openapi.Actor_Context
	if err := ApContext.FromActorContext1([]string{"https://www.w3.org/ns/activitystreams", "https://w3id.org/security/v1"}); err != nil {
		panic(err)
	}
	return ApContext
}

var ApContext = NewApContext()

var Protocols = []string{
	"activitypub",
}

const SoftWareName = "sechack365"
const SoftWareVersion = "0.1.0"

const NodeInfoVersion = "2.0"

// NOTE: const values: end

func (s *ActivitypubService) NewActorUrl(host, id string) string {
	return fmt.Sprintf("https://%s/api/v1/users/%s", host, id)
}

func (s *ActivitypubService) NewKeyIdUrl(host string, name string) string {
	return s.NewActorUrl(host, name) + "#main-key"
}

func (s *ActivitypubService) NewNodeInfo(userUsage int) *openapi.Nodeinfo {
	return &openapi.Nodeinfo{
		OpenRegistrations: false,
		Protocols:         Protocols,
		Software: openapi.NodeinfoSoftware{
			Name:    SoftWareName,
			Version: SoftWareVersion,
		},
		Usage: openapi.NodeinfoUsage{
			Users: openapi.NodeinfoUsageUsers{
				Total: 1,
			},
		},
		Services: openapi.NodeinfoService{
			Inbound:  map[string]interface{}{},
			Outbound: map[string]interface{}{},
		},
		Metadata: openapi.NodeinfoMetadata{},
		Version:  NodeInfoVersion,
	}
}

func (s *ActivitypubService) ResolveWebfinger(username, host string) (string, error) {
	url := fmt.Sprintf("https://%s/.well-known/webfinger?resource=acct:%s@%s", host, username, host)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", cerror.Wrap(cerror.ErrResolveWebfinger, err.Error())
	}

	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", cerror.Wrap(cerror.ErrResolveWebfinger, err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", cerror.Wrap(cerror.ErrResolveWebfinger, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return "", cerror.Wrap(cerror.ErrResolveWebfinger, string(body))
	}

	var data openapi.WellknownWebfinger
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", cerror.Wrap(cerror.ErrResolveWebfinger, err.Error())
	}

	var link string
	for _, l := range data.Links {
		if l.Rel == "self" {
			link = l.Href
			break
		}
	}

	if link == "" {
		return "", cerror.Wrap(cerror.ErrResolveWebfinger, "link not found")
	}

	return link, nil
}

func (s *ActivitypubService) ResolveRemoteActor(link string) (*openapi.Actor, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, cerror.Wrap(cerror.ErrResolveRemoteActor, err.Error())
	}

	req.Header.Set("Accept", "application/activity+json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, cerror.Wrap(cerror.ErrResolveRemoteActor, err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, cerror.Wrap(cerror.ErrResolveRemoteActor, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, cerror.Wrap(cerror.ErrResolveRemoteActor, string(body))
	}

	var actor openapi.Actor
	if err := json.Unmarshal(body, &actor); err != nil {
		return nil, cerror.Wrap(cerror.ErrResolveRemoteActor, err.Error())
	}

	return &actor, nil
}
