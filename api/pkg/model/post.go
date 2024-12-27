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
	User      UserInPost `db:"user"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}

type UserInPost struct {
	Username    string `db:"username"`
	DisplayName string `db:"display_name"`
	Icon        string `db:"icon"`
}
