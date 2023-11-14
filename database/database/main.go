package main

import (
	"github.com/JngMkk/golang/database/database/databasepkg"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := databasepkg.Setup()
	if err != nil {
		panic(err)
	}

	if err := databasepkg.Exec(db); err != nil {
		panic(err)
	}
}
