package repository

import "github.com/claustra01/sechack365/pkg/model"

type TransactionRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *TransactionRepository) Begin() error {
	_, err := r.SqlHandler.Execute("BEGIN;")
	return err
}

func (r *TransactionRepository) Commit() error {
	_, err := r.SqlHandler.Execute("COMMIT;")
	return err
}

func (r *TransactionRepository) Rollback() error {
	_, err := r.SqlHandler.Execute("ROLLBACK;")
	return err
}
