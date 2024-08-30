package model

type Follow struct {
	Follower   string `json:"follower"`
	Followee   string `json:"followee"`
	IsAccepted bool   `json:"is_accepted"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
