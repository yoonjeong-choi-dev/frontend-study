package redis

import (
	redis "gopkg.in/redis.v5"
	"os"
)

func Setup() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	// Check the connection
	_, err := client.Ping().Result()
	return client, err
}
