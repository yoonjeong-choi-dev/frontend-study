package transaction

import (
	"database/sql"
)

// DB Wrapping Interface for sql.DB
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Transaction Generic Interface for wrapping sql.Db & sql.Tx
type Transaction interface {
	DB
	Commit() error
	Rollback() error
}
