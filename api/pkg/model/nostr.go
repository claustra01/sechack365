package model

type NostrEvent struct {
	Id        string          `json:"id" db:"id"`
	Pubkey    string          `json:"pubkey" db:"pubkey"`
	CreatedAt int             `json:"created_at" db:"created_at"` // unix timestamp
	Kind      int             `json:"kind" db:"kind"`
	Tags      []NostrEventTag `json:"tags" db:"tags"`
	Content   string          `json:"content" db:"content"`
	Sig       string          `json:"sig" db:"sig"`
}

type NostrEventTag []string

type NostrFilter struct {
	// Ids     []string `json:"ids" db:"ids"`
	Authors []string `json:"authors" db:"authors"`
	Kinds   []int    `json:"kinds" db:"kinds"`
	// ETags   []string `json:"#e" db:"#e"`
	// PTags   []string `json:"#p" db:"#p"`
	Since int64 `json:"since" db:"since"` // unix timestamp
	Until int64 `json:"until" db:"until"` // unix timestamp
	Limit int   `json:"limit" db:"limit"`
}

type NostrProfile struct {
	Name        string `json:"name" db:"name"`
	DisplayName string `json:"display_name" db:"display_name"`
	About       string `json:"about" db:"about"`
	Picture     string `json:"picture" db:"picture"`
}
