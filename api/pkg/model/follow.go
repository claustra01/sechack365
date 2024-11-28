package model

type Follow struct {
	Id         string `json:"id" db:"id"`
	FollowerId string `json:"follower_id" db:"follower_id"`
	FolloweeId string `json:"followee_id" db:"followee_id"`
	IsAccepted bool   `json:"is_accepted" db:"is_accepted"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
}
