package usecase

import "github.com/claustra01/sechack365/pkg/model"

type ITransactionRepository interface {
	Begin() (model.Tx, error)
}

type TransactionUsecase struct {
	TransactionRepository ITransactionRepository
}

func (u *TransactionUsecase) Begin() (model.Tx, error) {
	return u.TransactionRepository.Begin()
}
