package model

import (
	"context"
	"io"
)

type IStorageHandler interface {
	GetObject(ctx context.Context, bucket string, objKey string) (io.Reader, error)
	PutObject(ctx context.Context, bucket string, objKey string, reader io.Reader, size int64) error
	RemoveObject(ctx context.Context, bucket string, objKey string) error
}
