package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type SqlHandler struct {
	Pg *sql.DB
}

func Connect(connStr string) (*SqlHandler, error) {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &SqlHandler{Pg: conn}, nil
}

func Close(conn *SqlHandler) error {
	return conn.Pg.Close()
}
