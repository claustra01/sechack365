package model

type ISqlHandler interface {
	Execute(string, ...any) (Result, error)
	Query(string, ...any) (Row, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...any) error
	Next() bool
	Close() error
}
