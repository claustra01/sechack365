package model

import "time"

type Post struct {
	Id        string    `db:"id"`
	UserId    string    `db:"user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type PostWithUser struct {
	Id        string     `db:"id"`
	UserId    string     `db:"user_id"`
	Content   string     `db:"content"`
	User      SimpleUser `db:"user"`
	LikeCount int        `db:"like_count"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}
