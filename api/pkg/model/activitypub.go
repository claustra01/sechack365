package model

import "time"

const (
	SoftWareName    = "sechack365"
	SoftWareVersion = "0.1.0"
	NodeInfoVersion = "2.0"
)

const (
	ActivityTypePerson = "Person"
	ActivityTypeCreate = "Create"
	ActivityTypeUndo   = "Undo"
	ActivityTypeFollow = "Follow"
	ActivityTypeNote   = "Note"
	ActivityTypeAccept = "Accept"
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

type ApNoteActivity struct {
	Type      string    `json:"type"`
	Id        string    `json:"id"`
	Content   string    `json:"content"`
	Published time.Time `json:"published"`
	To        []string  `json:"to"` // NOTE: まだ使わなくていい
	Cc        []string  `json:"cc"` // NOTE: まだ使わなくていい
}
