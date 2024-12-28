package model

import "time"

type Reaction struct {
	Id        string    `db:"id"`
	Type      string    `db:"type"`
	UserId    string    `db:"user_id"`
	PostId    string    `db:"post_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
