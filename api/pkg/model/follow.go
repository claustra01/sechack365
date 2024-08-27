package model

type Follow struct {
	Id         string `json:"id"`
	Follower   string `json:"follower"`
	Followee   string `json:"followee"`
	IsAccepted bool   `json:"is_accepted"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
