package model

type ISqlHandler interface {
	Exec(string, ...any) (Result, error)
	Query(string, ...any) (Row, error)
	Select(any, string, ...any) error
	Get(any, string, ...any) error
	Begin() (Tx, error)
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

type Tx interface {
	Commit() error
	Rollback() error
}
