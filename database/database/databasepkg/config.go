package databasepkg

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Example struct {
	Name    string
	Created *time.Time
}

// 데이터베이스를 설정하고 연결 풀이 설정된 데이터베이스를 반환
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/gocookbook?parseTime=true", os.Getenv("MYSQLUSERNAME"), os.Getenv("MYSQLPASSWORD")))
	if err != nil {
		return nil, err
	}
	return db, nil
}
