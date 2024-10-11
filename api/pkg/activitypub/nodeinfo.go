package activitypub

import "github.com/claustra01/sechack365/pkg/openapi"

// TODO: usage should be included users, posts, and more
func BuildNodeInfoSchema(usersUsage int) *openapi.Nodeinfo {
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
