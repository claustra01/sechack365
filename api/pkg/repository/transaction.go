package repository

import "github.com/claustra01/sechack365/pkg/model"

type TransactionRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *TransactionRepository) Begin() (model.Tx, error) {
	return r.SqlHandler.Begin()
}
