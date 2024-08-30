package activitypub

type NodeInfo struct {
	OpenRegistrations bool     `json:"openRegistrations"`
	Protocols         []string `json:"protocols"`
	Software          Software `json:"software"`
	Usage             Usage    `json:"usage"`
	Services          Services `json:"services"`
	Metadata          Metadata `json:"metadata"`
	Version           string   `json:"version"`
}

type Software struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Usage struct {
	Users UsersUsage `json:"users"`
}

type UsersUsage struct {
	Total int `json:"total"`
}

type Services struct {
	Inbound  struct{} `json:"inbound"`
	Outbound struct{} `json:"outbound"`
}

type Metadata struct{}

// TODO: usage should be included users, posts, and more
func BuildNodeInfoSchema(usersUsage int) *NodeInfo {
	return &NodeInfo{
		OpenRegistrations: false,
		Protocols:         Protocols,
		Software: Software{
			Name:    SoftWareName,
			Version: SoftWareVersion,
		},
		Usage: Usage{
			Users: UsersUsage{
				Total: 1,
			},
		},
		Services: Services{
			Inbound:  struct{}{},
			Outbound: struct{}{},
		},
		Metadata: Metadata{},
		Version:  NodeInfoVersion,
	}
}
