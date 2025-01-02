package model

const SoftWareName = "sechack365"
const SoftWareVersion = "0.1.0"

const NodeInfoVersion = "2.0"

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
	Object  string `json:"object"`
}
