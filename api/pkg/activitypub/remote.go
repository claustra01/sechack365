package activitypub

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/openapi"
)

func ResolveWebfinger(username string, host string) (string, error) {
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

func ResolveRemoteActor(link string) (*openapi.Actor, error) {
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
