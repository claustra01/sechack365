package repository

import "github.com/claustra01/sechack365/pkg/model"

type TransactionRepository struct {
	SqlHandler model.ISqlHandler
}

func (repo *TransactionRepository) Begin() error {
	_, err := repo.SqlHandler.Execute("BEGIN;")
	return err
}

func (repo *TransactionRepository) Commit() error {
	_, err := repo.SqlHandler.Execute("COMMIT;")
	return err
}

func (repo *TransactionRepository) Rollback() error {
	_, err := repo.SqlHandler.Execute("ROLLBACK;")
	return err
}
