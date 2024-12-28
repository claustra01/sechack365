package model

import "time"

type Follow struct {
	Id         string    `db:"id"`
	FollowerId string    `db:"follower_id"`
	TargetId   string    `db:"target_id"`
	IsAccepted bool      `db:"is_accepted"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
