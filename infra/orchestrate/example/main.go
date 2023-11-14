package main

import "github.com/JngMkk/golang/infra/orchestrate"

func main() {
	if err := orchestrate.Exec("mongodb://mongodb:27017"); err != nil {
		panic(err)
	}
}
