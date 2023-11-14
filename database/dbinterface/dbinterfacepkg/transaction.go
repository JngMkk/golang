package dbinterfacepkg

import "database/sql"

// sql.DB 또는 sql.Transaction을 만족하는 인터페이스
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// commit, rollback, Stmt 등 쿼리가 할 수 잇는 모든 작업을 할 수 있음
type Transaction interface {
	DB
	Commit() error
	Rollback() error
}
