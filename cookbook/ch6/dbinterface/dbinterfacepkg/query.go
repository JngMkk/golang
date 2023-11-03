package dbinterfacepkg

import (
	"fmt"

	"github.com/JngMkk/cookbook/ch6/database/databasepkg"
)

// 새로운 연결을 포착해 테이블을 생성하고 몇 가지 쿼리를 수행한 다음, 테이블을 제거
func Query(db DB) error {
	name := "kim"
	rows, err := db.Query("SELECT name, created FROM example WHERE name=?", name)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var e databasepkg.Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("Results:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}

	return rows.Err()
}
