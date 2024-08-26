package util

import "github.com/google/uuid"

func NewUuid() uuid.UUID {
	return uuid.Must(uuid.NewV7())
}
