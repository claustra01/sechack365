package model

import "time"

type NostrRelay struct {
	Id        string    `db:"id"`
	Url       string    `db:"url"`
	IsEnable  bool      `db:"is_enable"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
