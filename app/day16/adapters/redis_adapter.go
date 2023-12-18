package adapters

import "github.com/go-redis/redis"


type  RedisAdapter struct {
	client *redis.Client
}

func NewRedisAdapter() *RedisAdapter {
	return &RedisAdapter{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6381",
			Password: "",
			DB:       0,
		}),
	}
}

func (adapter *RedisAdapter) Get(key string) string {
	val, err := adapter.client.Get(key).Result()
	if err != nil {
		return ""
	}
	return val
}

func (adapter * RedisAdapter) Set(key string, value string) {
	adapter.client.Set(key, value, 0)
}

func (adapter * RedisAdapter) Delete(key string) {
	adapter.client.Del(key)
}