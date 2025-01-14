package infrastructure

import (
	"context"
	"fmt"
	"io"

	"github.com/claustra01/sechack365/pkg/model"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type StorageHandler struct {
	Client *minio.Client
}

func NewStorageHandler(host, port, key, secret string) (model.IStorageHandler, error) {
	client, err := minio.New(fmt.Sprintf("%s:%s", host, port), &minio.Options{
		Creds:  credentials.NewStaticV4(key, secret, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}
	storageHandler := new(StorageHandler)
	storageHandler.Client = client
	return storageHandler, nil
}

func (conn *StorageHandler) GetObject(ctx context.Context, bucket string, objKey string) (io.Reader, error) {
	return conn.Client.GetObject(ctx, bucket, objKey, minio.GetObjectOptions{})
}

func (conn *StorageHandler) PutObject(ctx context.Context, bucket string, objKey string, reader io.Reader, size int64) error {
	_, err := conn.Client.PutObject(ctx, bucket, objKey, reader, size, minio.PutObjectOptions{})
	return err
}

func (conn *StorageHandler) RemoveObject(ctx context.Context, bucket string, objKey string) error {
	return conn.Client.RemoveObject(ctx, bucket, objKey, minio.RemoveObjectOptions{})
}
