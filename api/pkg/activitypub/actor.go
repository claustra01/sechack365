package activitypub

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/openapi"
)

func BuildActorSchema(user model.User, identifier model.ApUserIdentifier) *openapi.Actor {
	baseUrl := BuildActorUrl(user.Host, user.Id)
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
			Id:           BuildKeyIdUrl(user.Host, user.Username),
			Owner:        baseUrl,
			PublicKeyPem: identifier.PublicKey,
		},
	}
	return actor
}
