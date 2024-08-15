package activitypub

import (
	"fmt"
	"log/slog"

	"github.com/claustra01/sechack365/pkg/util"
)

type Actor struct {
	Context           []string       `json:"@context"`
	Type              string         `json:"type"`
	Id                string         `json:"id"`
	Inbox             string         `json:"inbox"`
	Outbox            string         `json:"outbox"`
	PreferredUsername string         `json:"preferredUsername"`
	Name              string         `json:"name"`
	Summary           string         `json:"summary"`
	PublicKey         util.PublicKey `json:"publicKey"`
}

func GetActor(name string, host string) (*Actor, error) {
	baseUrl := fmt.Sprintf("https://%s/actor/%s", host, name)

	publicKey, _, err := util.GenerateKeyPair()
	if err != nil {
		slog.Error("PublicKey Generation Error:", "Error", err)
		return nil, err
	}

	actor := &Actor{
		Context: []string{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
		},
		Type:              "Person",
		Id:                baseUrl,
		Inbox:             baseUrl + "/inbox",
		Outbox:            baseUrl + "/outbox",
		PreferredUsername: "mock",
		Name:              "Mock User",
		Summary:           "The user in activitypub server made by claustra01",
		PublicKey: util.PublicKey{
			Type:         "Key",
			Id:           baseUrl + "#main-key",
			Owner:        baseUrl,
			PublicKeyPem: publicKey,
		},
	}

	return actor, nil
}
