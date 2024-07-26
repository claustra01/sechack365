package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/claustra01/sechack365/pkg/activitypub"
	"github.com/claustra01/sechack365/pkg/framework"
)

func resolveWebfinger(u string, h string) (string, error) {
	url := fmt.Sprintf("https://%s/.well-known/webfinger?resource=acct:%s@%s", h, u, h)
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
		return "", fmt.Errorf("status code is not 200")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data activitypub.Webfinger
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

func resolveRemoteActor(link string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/activity+json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code is not 200")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func LookupActor(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		user := query.Get("user")

		re := regexp.MustCompile(`^([a-zA-Z0-9_]+)@?([a-zA-Z0-9.]+)?$`)
		matches := re.FindStringSubmatch(user)
		log.Print(user, matches)

		if len(matches) != 3 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		u := matches[1]
		h := matches[2]

		if h == c.Config.Host || h == "" {
			http.Redirect(w, r, "/api/v1/actor/"+u, http.StatusSeeOther)
			return
		}

		link, err := resolveWebfinger(u, h)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		actor, err := resolveRemoteActor(link)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, actor)
	}
}
