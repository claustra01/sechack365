package controller

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/repository"
	"github.com/claustra01/sechack365/pkg/usecase"
)

type FileController struct {
	FileUsecase usecase.FileUsecase
}

func NewFileController(conn model.IStorageHandler) *FileController {
	return &FileController{
		FileUsecase: usecase.FileUsecase{
			FileRepository: &repository.FileRepository{
				StorageHandler: conn,
			},
		},
	}
}

func (c *FileController) SaveImage(file []byte, filename string, bucket string) error {
	return c.FileUsecase.SaveImage(file, filename, bucket)
}
