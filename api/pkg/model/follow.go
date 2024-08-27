package model

type Follow struct {
	Id        string `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
