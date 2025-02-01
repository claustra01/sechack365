package model

import "time"

type Article struct {
	Id        string    `db:"id"`
	UserId    string    `db:"user_id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ArticleComment struct {
	Id        string    `db:"id"`
	UserId    string    `db:"user_id"`
	ArticleId string    `db:"article_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ArticleWithUser struct {
	Id        string     `db:"id"`
	Title     string     `db:"title"`
	Content   string     `db:"content"`
	User      SimpleUser `db:"user"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}

type ArticleCommentWithUser struct {
	Id        string     `db:"id"`
	ArticleId string     `db:"article_id"`
	Content   string     `db:"content"`
	User      SimpleUser `db:"user"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}

type ArticlePostRelation struct {
	ArticleId string `db:"article_id"`
	PostId    string `db:"post_id"`
}
