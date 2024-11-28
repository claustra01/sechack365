package infrastructure

import (
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SqlHandler struct {
	Pg *sqlx.DB
}

type Tx struct {
	Tx *sqlx.Tx
}

func NewSqlHandler(connStr string) (model.ISqlHandler, error) {
	conn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Pg = conn
	return sqlHandler, nil
}

func (conn *SqlHandler) Close() error {
	return conn.Pg.Close()
}

func (conn *SqlHandler) Exec(query string, args ...interface{}) (model.Result, error) {
	result, err := conn.Pg.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (conn *SqlHandler) Query(query string, args ...any) (model.Row, error) {
	return conn.Pg.Query(query, args...)
}

func (conn *SqlHandler) Select(dest any, query string, args ...any) error {
	return conn.Pg.Select(dest, query, args...)
}

func (conn *SqlHandler) Get(dest any, query string, args ...any) error {
	return conn.Pg.Get(dest, query, args...)
}

func (conn *SqlHandler) Begin() (model.Tx, error) {
	tx, err := conn.Pg.Beginx()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (tx *Tx) Commit() error {
	return tx.Tx.Commit()
}

func (tx *Tx) Rollback() error {
	return tx.Tx.Rollback()
}
