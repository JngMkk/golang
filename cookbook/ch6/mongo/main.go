package main

import "github.com/JngMkk/cookbook/ch6/mongo/mongodb"

func main() {
	if err := mongodb.Exec("mongodb://localhost"); err != nil {
		panic(err)
	}
}
