package redis

import (
	"gopkg.in/redis.v5"
	"time"
)

func SaveSingleData(client *redis.Client, key, value string) {
	client.Set(key, value, 5*time.Second)
}

func SaveListData(client *redis.Client, key string, values []int64) error {
	if len(values) == 0 {
		return nil
	}

	if err := client.LPush(key, values[0]).Err(); err != nil {
		return err
	}

	for i := 1; i < len(values); i++ {
		if err := client.LPush(key, values[i]).Err(); err != nil {
			// 실패 시 전체 리스트 제거 함수
			client.Del(key)
			return err
		}
	}
	return nil
}
