package main

import "github.com/JngMkk/cookbook/ch6/redisex/redispkg"

func main() {
	if err := redispkg.Exec(); err != nil {
		panic(err)
	}

	if err := redispkg.Sort(); err != nil {
		panic(err)
	}
}
