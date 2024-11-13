package model

type Post struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
