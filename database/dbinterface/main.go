package main

import (
	"github.com/JngMkk/golang/database/database/databasepkg"
	"github.com/JngMkk/golang/database/dbinterface/dbinterfacepkg"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := databasepkg.Setup()
	if err != nil {
		panic(err)
	}

	// transaction
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	if err := dbinterfacepkg.Exec(tx); err != nil {
		panic(err)
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}
}
