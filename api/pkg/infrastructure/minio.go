package infrastructure

import (
	"fmt"

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
