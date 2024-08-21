package infrastructure

import (
	"database/sql"

	"github.com/claustra01/sechack365/pkg/model"
	_ "github.com/lib/pq"
)

type SqlHandler struct {
	Pg *sql.DB
}

func NewSqlHandler(connStr string) (model.ISqlHandler, error) {
	conn, err := sql.Open("postgres", connStr)
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

func (conn *SqlHandler) Execute(query string, args ...any) (model.Result, error) {
	return conn.Pg.Exec(query, args...)
}

func (conn *SqlHandler) Query(query string, args ...any) (model.Row, error) {
	return conn.Pg.Query(query, args...)
}
