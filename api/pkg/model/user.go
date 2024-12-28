package model

import "time"

type User struct {
	Id             string    `db:"id"`
	Username       string    `db:"username"`
	Protocol       string    `db:"protocol"`
	HashedPassword string    `db:"hashed_password"`
	DisplayName    string    `db:"display_name"`
	Profile        string    `db:"profile"`
	Icon           string    `db:"icon"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type ApUserIdentifier struct {
	UserId        string    `db:"user_id"`
	LocalUsername string    `db:"local_username"`
	Host          string    `db:"host"`
	PublicKey     string    `db:"public_key"`
	PrivateKey    string    `db:"private_key"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type NostrUserIdentifier struct {
	UserId     string    `db:"user_id"`
	PublicKey  string    `db:"public_key"`
	PrivateKey string    `db:"private_key"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type UserWithIdentifiers struct {
	Id             string      `db:"id"`
	Username       string      `db:"username"`
	Protocol       string      `db:"protocol"`
	HashedPassword string      `db:"hashed_password"`
	DisplayName    string      `db:"display_name"`
	Profile        string      `db:"profile"`
	Icon           string      `db:"icon"`
	CreatedAt      time.Time   `db:"created_at"`
	UpdatedAt      time.Time   `db:"updated_at"`
	Identifiers    Identifiers `db:"identifiers"`
	PostCount      int         `db:"post_count"`
	FollowCount    int         `db:"follow_count"`
	FollowerCount  int         `db:"follower_count"`
}

type Identifiers struct {
	Activitypub *ApUserIdentifier    `db:"activitypub"`
	Nostr       *NostrUserIdentifier `db:"nostr"`
}

type SimpleUser struct {
	Username    string `db:"username"`
	Protocol    string `db:"protocol"`
	DisplayName string `db:"display_name"`
	Icon        string `db:"icon"`
}
