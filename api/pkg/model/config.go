package model

type NostrRelay struct {
	Id        string `json:"id" db:"id"`
	Url       string `json:"url" db:"url"`
	IsEnable  bool   `json:"is_enable" db:"is_enable"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
