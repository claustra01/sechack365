package activitypub

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/claustra01/sechack365/pkg/cerror"
)

func ResolveWebfinger(username string, host string) (string, error) {
	url := fmt.Sprintf("https://%s/.well-known/webfinger?resource=acct:%s@%s", host, username, host)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", cerror.ErrResolveWebfinger
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data Webfinger
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	var link string
	for _, l := range data.Links {
		if l.Rel == "self" {
			link = l.Href
			break
		}
	}

	if link == "" {
		return "", fmt.Errorf("link not found")
	}

	return link, nil
}

func ResolveRemoteActor(link string) (*Actor, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/activity+json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, cerror.ErrResolveRemoteActor
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var actor Actor
	if err := json.Unmarshal(body, &actor); err != nil {
		return nil, err
	}

	return &actor, nil
}
