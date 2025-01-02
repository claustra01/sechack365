package service

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/openapi"
	"github.com/claustra01/sechack365/pkg/util"
)

type ActivitypubService struct{}

func NewApContext() openapi.Actor_Context {
	var ApContext openapi.Actor_Context
	if err := ApContext.FromActorContext1(model.ApContext); err != nil {
		panic(err)
	}
	return ApContext
}

func (s *ActivitypubService) NewNodeInfo(userUsage int) *openapi.Nodeinfo {
	return &openapi.Nodeinfo{
		OpenRegistrations: false,
		Protocols:         model.Protocols,
		Software: openapi.NodeinfoSoftware{
			Name:    model.SoftWareName,
			Version: model.SoftWareVersion,
		},
		Usage: openapi.NodeinfoUsage{
			Users: openapi.NodeinfoUsageUsers{
				Total: userUsage,
			},
		},
		Services: openapi.NodeinfoService{
			Inbound:  map[string]interface{}{},
			Outbound: map[string]interface{}{},
		},
		Metadata: openapi.NodeinfoMetadata{},
		Version:  model.NodeInfoVersion,
	}
}

func (s *ActivitypubService) NewActor(user model.UserWithIdentifiers) *openapi.Actor {
	baseUrl := s.NewActorUrl(user.Identifiers.Activitypub.Host, user.Id)
	actor := &openapi.Actor{
		Context:           NewApContext(),
		Type:              "Person",
		Id:                baseUrl,
		Inbox:             baseUrl + "/inbox",
		Outbox:            baseUrl + "/outbox",
		PreferredUsername: user.Username,
		Name:              user.DisplayName,
		Summary:           user.Profile,
		Icon: openapi.ActorIcon{
			Type: "Image",
			Url:  user.Icon,
		},
		PublicKey: openapi.ActorPublicKey{
			Type:         "Key",
			Id:           s.NewActorUrl(user.Identifiers.Activitypub.Host, user.Username),
			Owner:        baseUrl,
			PublicKeyPem: user.Identifiers.Activitypub.PublicKey,
		},
	}
	return actor
}

func (s *ActivitypubService) NewActorUrl(host, id string) string {
	return fmt.Sprintf("https://%s/api/v1/users/%s", host, id)
}

func (s *ActivitypubService) NewKeyIdUrl(host string, name string) string {
	return s.NewActorUrl(host, name) + "#main-key"
}

func (s *ActivitypubService) NewFollowActivity(id, host, followerId, targetUrl string) *model.ApActivity {
	return &model.ApActivity{
		Context: NewApContext(),
		Type:    "Follow",
		Id:      fmt.Sprintf("https://%s/follows/%s", host, id),
		Actor:   s.NewActorUrl(host, followerId),
		Object:  targetUrl,
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

func (s *ActivitypubService) SendActivity(keyId string, privKey *rsa.PrivateKey, targetHost string, activity any) ([]byte, error) {
	reqBody, err := json.Marshal(activity)
	if err != nil {
		return nil, cerror.Wrap(err, "failed to send activity")
	}

	// create request
	client := &http.Client{}
	req, err := http.NewRequest("POST", targetHost, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, cerror.Wrap(err, "failed to send activity")
	}
	req.Header.Set("Content-Type", "application/activity+json")
	if err := util.HttpSigSign(keyId, privKey, req, reqBody); err != nil {
		return nil, cerror.Wrap(err, "failed to send activity")
	}

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, cerror.Wrap(err, "failed to send activity")
	}
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, cerror.Wrap(err, "failed to send activity")
	}
	fmt.Println("status code: ", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return nil, cerror.Wrap(fmt.Errorf("%v", string(body)), "failed to send activity")
	}
	return body, nil
}
