package activitypub

import (
	"fmt"

	"github.com/claustra01/sechack365/pkg/model"
)

type PublicKey struct {
	Type         string `json:"type"`
	Id           string `json:"id"`
	Owner        string `json:"owner"`
	PublicKeyPem string `json:"publicKeyPem"`
}

type Actor struct {
	Context           []string  `json:"@context"`
	Type              string    `json:"type"`
	Id                string    `json:"id"`
	Inbox             string    `json:"inbox"`
	Outbox            string    `json:"outbox"`
	PreferredUsername string    `json:"preferredUsername"`
	Name              string    `json:"name"`
	Summary           string    `json:"summary"`
	PublicKey         PublicKey `json:"publicKey"`
}

func GetActor(user model.ApUser) *Actor {
	baseUrl := fmt.Sprintf("https://%s/actor/%s", user.Host, user.Username)
	actor := &Actor{
		Context:           ApContext[:],
		Type:              "Person",
		Id:                baseUrl,
		Inbox:             baseUrl + "/inbox",
		Outbox:            baseUrl + "/outbox",
		PreferredUsername: user.Username,
		Name:              user.DisplayName,
		Summary:           user.Profile,
		PublicKey: PublicKey{
			Type:         "Key",
			Id:           baseUrl + "#main-key",
			Owner:        baseUrl,
			PublicKeyPem: user.PublicKey,
		},
	}
	return actor
}
