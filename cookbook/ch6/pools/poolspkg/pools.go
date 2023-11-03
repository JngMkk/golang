package poolspkg

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// pools 패키지를 사용하고 연결 수 제한 등을 적용해 db를 구성함
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/gocookbook?parseTime=true", os.Getenv("MYSQLUSERNAME"), os.Getenv("MYSQLPASSWORD")))
	if err != nil {
		return nil, err
	}

	// 최대 24개의 연결만 허용
	db.SetMaxOpenConns(24)

	// MaxIdleConns는 열려 있는 최대 SetMaxOpenConns보다 작을 수 없음
	// 그렇지 않으면 해당 값으로 기본 설정됨
	db.SetMaxIdleConns(24)

	return db, nil
}
