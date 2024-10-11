package activitypub

import (
	"github.com/claustra01/sechack365/pkg/openapi"
)

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
