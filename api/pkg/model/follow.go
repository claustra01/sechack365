package model

type Follow struct {
	Id        string `json:"id"`
	Follower  string `json:"follower"`
	Followee  string `json:"followee"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
