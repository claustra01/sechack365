package activitypub

import (
	"encoding/json"

	"github.com/claustra01/sechack365/pkg/openapi"
)

var ApContext = openapi.Actor_Context{
	union: json.RawMessage(`["https://www.w3.org/ns/activitystreams","https://w3id.org/security/v1"]`),
}

var Protocols = []string{
	"activitypub",
}

const SoftWareName = "sechack365"
const SoftWareVersion = "0.1.0"

const NodeInfoVersion = "2.0"
