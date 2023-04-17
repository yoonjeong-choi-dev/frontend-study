package main

import (
	"github.com/joho/godotenv"
	"redis"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	if err := redis.RedisExample(); err != nil {
		panic(err)
	}
}
