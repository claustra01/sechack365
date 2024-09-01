package model

type Follow struct {
	Id         string `json:"id"`
	FollowerId string `json:"follower_id"`
	FolloweeId string `json:"followee_id"`
	IsAccepted bool   `json:"is_accepted"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
