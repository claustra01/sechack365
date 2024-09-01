package controller

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/repository"
	"github.com/claustra01/sechack365/pkg/usecase"
)

type TransactionController struct {
	TransactionUsecase usecase.TransactionUsecase
}

func NewTransactionController(conn model.ISqlHandler) *TransactionController {
	return &TransactionController{
		TransactionUsecase: usecase.TransactionUsecase{
			TransactionRepository: &repository.TransactionRepository{
				SqlHandler: conn,
			},
		},
	}
}

func (c *TransactionController) Begin() error {
	return c.TransactionUsecase.Begin()
}

func (c *TransactionController) Commit() error {
	return c.TransactionUsecase.Commit()
}

func (c *TransactionController) Rollback() error {
	return c.TransactionUsecase.Rollback()
}
