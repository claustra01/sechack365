package activitypub

import (
	"github.com/claustra01/sechack365/pkg/model"
)

type Image struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type PublicKey struct {
	Type         string `json:"type"`
	Id           string `json:"id"`
	Owner        string `json:"owner"`
	PublicKeyPem string `json:"publicKeyPem"`
}

type Actor struct {
	Context           any       `json:"@context"`
	Type              string    `json:"type"`
	Id                string    `json:"id"`
	Inbox             string    `json:"inbox"`
	Outbox            string    `json:"outbox"`
	PreferredUsername string    `json:"preferredUsername"`
	Name              string    `json:"name"`
	Summary           string    `json:"summary"`
	Icon              Image     `json:"icon"`
	PublicKey         PublicKey `json:"publicKey"`
}

func BuildActorSchema(user model.ApUser) *Actor {
	baseUrl := BuildActorUrl(user.Host, user.Username)
	actor := &Actor{
		Context:           ApContext[:],
		Type:              "Person",
		Id:                baseUrl,
		Inbox:             baseUrl + "/inbox",
		Outbox:            baseUrl + "/outbox",
		PreferredUsername: user.Username,
		Name:              user.DisplayName,
		Summary:           user.Profile,
		Icon: Image{
			Type: "Image",
			Url:  user.Icon,
		},
		PublicKey: PublicKey{
			Type:         "Key",
			Id:           baseUrl + "#main-key",
			Owner:        baseUrl,
			PublicKeyPem: user.PublicKey,
		},
	}
	return actor
}
