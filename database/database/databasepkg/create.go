package databasepkg

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Create(db *sql.DB) error {
	// 데이터베이스 생성
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO example (name, created) values ("kim", NOW())`); err != nil {
		return err
	}

	return nil
}
