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

func (controller *TransactionController) Begin() error {
	return controller.TransactionUsecase.Begin()
}

func (controller *TransactionController) Commit() error {
	return controller.TransactionUsecase.Commit()
}

func (controller *TransactionController) Rollback() error {
	return controller.TransactionUsecase.Rollback()
}
