package main

import "github.com/JngMkk/cookbook/ch11/orchestrate"

func main() {
	if err := orchestrate.Exec("mongodb://mongodb:27017"); err != nil {
		panic(err)
	}
}
