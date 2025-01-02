package model

const (
	SoftWareName    = "sechack365"
	SoftWareVersion = "0.1.0"
	NodeInfoVersion = "2.0"
)

const (
	ActivityTypePerson = "Person"
	ActivityTypeFollow = "Follow"
	ActivityTypeAccept = "Accept"
	ActivityTypeUndo   = "Undo"
	ActivityTypeReject = "Reject"
)

var Protocols = []string{
	"activitypub",
}

var ApContext = []string{
	"https://www.w3.org/ns/activitystreams",
	"https://w3id.org/security/v1",
}

type ApActivity struct {
	Context any    `json:"@context"`
	Type    string `json:"type"`
	Id      string `json:"id"`
	Actor   string `json:"actor"`
	Object  any    `json:"object"`
}
