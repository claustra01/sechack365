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

func (c *TransactionController) Begin() (model.Tx, error) {
	return c.TransactionUsecase.Begin()
}
