package redis

import (
	"fmt"
	"gopkg.in/redis.v5"
)

func QuerySingleValue(client *redis.Client, key string) error {
	var result string

	if err := client.Get(key).Scan(&result); err != nil {
		switch err {
		case redis.Nil:
			result = "NO DATA"
		default:
			return err
		}
	}
	fmt.Printf("Query with key '%s': %s\n", key, result)
	return nil
}

func QueryListValueBySort(client *redis.Client, key string) error {
	res, err := client.Sort(key, redis.Sort{Order: "ASC"}).Result()
	if err != nil {
		return err
	}
	fmt.Printf("List Data with key '%s': %v\n", key, res)
	return nil
}
