package repository

import (
	"bytes"
	"context"

	"github.com/claustra01/sechack365/pkg/model"
)

type FileRepository struct {
	StorageHandler model.IStorageHandler
}

func (r *FileRepository) SaveImage(file []byte, filename string, bucket string) error {
	if err := r.StorageHandler.PutObject(context.Background(), bucket, filename, bytes.NewReader(file), int64(len(file))); err != nil {
		return err
	}
	return nil
}
