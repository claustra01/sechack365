package service

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/openapi"
	"github.com/claustra01/sechack365/pkg/usecase"
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

func (s *ActivitypubService) NewActor(user model.User, identifier model.ApUserIdentifier) *openapi.Actor {
	baseUrl := s.NewActorUrl(user.Host, user.Id)
	actor := &openapi.Actor{
		Context:           ApContext,
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
			Id:           s.NewActorUrl(user.Host, user.Username),
			Owner:        baseUrl,
			PublicKeyPem: identifier.PublicKey,
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

func (s *ActivitypubService) NewFollowActivity(id, host, followerId, followeeUrl string) *usecase.FollowActivity {
	// object is followee actor
	return &usecase.FollowActivity{
		Context: ApContext,
		Type:    "Follow",
		Id:      fmt.Sprintf("https://%s/follows/%s", host, id),
		Actor:   s.NewActorUrl(host, followerId),
		Object:  followeeUrl,
	}
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

func (s *ActivitypubService) SendActivity(url string, activity any, host string, keyId string, prvKey *rsa.PrivateKey) ([]byte, error) {
	reqBody, err := json.Marshal(activity)
	if err != nil {
		return nil, cerror.Wrap(cerror.ErrPushActivity, err.Error())
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, cerror.Wrap(cerror.ErrPushActivity, err.Error())
	}

	signedDate := time.Now().Format(http.TimeFormat)

	req.Header.Set("Host", host)
	req.Header.Set("Date", signedDate)
	req.Header.Set("Content-Type", "application/activity+json")

	hash := sha256.Sum256(reqBody)
	digest := base64.StdEncoding.EncodeToString(hash[:])
	digestHeader := fmt.Sprintf("SHA-256=%s", digest)
	req.Header.Set("Digest", digestHeader)

	// TODO: このsigningStringを署名するのが正しいという認識だが理解が怪しいので後日実装する
	// signingString := fmt.Sprintf("(request-target): post %s\nhost: %s\ndate: %s\ndigest: %s", url, sigParams.Host, signedDate, digestHeader)
	rawSign, err := rsa.SignPKCS1v15(rand.Reader, prvKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, cerror.Wrap(cerror.ErrPushActivity, err.Error())
	}
	encodedSign := base64.StdEncoding.EncodeToString(rawSign)
	signatureHeader := fmt.Sprintf(`keyId="%s",algorithm="rsa-sha256",headers="(request-target) host date digest",signature="%s"`, keyId, encodedSign)
	req.Header.Set("Signature", signatureHeader)

	resp, err := client.Do(req)
	if err != nil {
		return nil, cerror.Wrap(cerror.ErrPushActivity, err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, cerror.Wrap(cerror.ErrPushActivity, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, cerror.Wrap(cerror.ErrPushActivity, string(body))
	}
	return body, nil
}
