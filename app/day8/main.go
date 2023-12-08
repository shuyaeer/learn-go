package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

type RedisAdapter struct {
	client *redis.Client
}

func NewRedisAdapter() *RedisAdapter {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // コンテナの名前:ポート番号
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return &RedisAdapter{
		client: client,
	}
}

func (r *RedisAdapter) Set(key string, value string) error {
	err := r.client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisAdapter) Get(key string) (string, error) {
	val, err := r.client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func main() {
	redisAdapter := NewRedisAdapter()

	err := redisAdapter.Set("key", "valueです")
	if err != nil {
		fmt.Println("Error setting value:", err)
	}

	value, err := redisAdapter.Get("key")
	if err != nil {
		fmt.Println("Error getting value:", err)
	} else {
		fmt.Println("Value:", value)
	}
}


