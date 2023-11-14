package main

import "github.com/JngMkk/golang/database/mongo/mongodb"

func main() {
	if err := mongodb.Exec("mongodb://localhost"); err != nil {
		panic(err)
	}
}
