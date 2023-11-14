package main

import "github.com/JngMkk/golang/database/redis/redispkg"

func main() {
	if err := redispkg.Exec(); err != nil {
		panic(err)
	}

	if err := redispkg.Sort(); err != nil {
		panic(err)
	}
}
