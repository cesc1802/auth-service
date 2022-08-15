package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(host, port, password string, db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})
	return &RedisClient{
		client: client,
	}
}

func (r *RedisClient) Keys(pattern string) ([]string, error) {
	return r.client.Keys(context.Background(), pattern).Result()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}

func (r *RedisClient) Set(key, value string) error {
	return r.client.Set(context.Background(), key, value, 0).Err()
}

func (r *RedisClient) Delete(key string) error {
	return r.client.Del(context.Background(), key).Err()
}

func (r *RedisClient) Pipeline() redis.Pipeliner {
	return r.client.Pipeline()
}
