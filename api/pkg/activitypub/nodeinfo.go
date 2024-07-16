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

func GetNodeInfo() *NodeInfo {
	// mock nodeinfo
	return &NodeInfo{
		OpenRegistrations: false,
		Protocols: []string{
			"activitypub",
		},
		Software: Software{
			Name:    "sechack365",
			Version: "0.1.0",
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
		Version:  "2.1",
	}
}
