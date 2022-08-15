package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/pkg/cache"
	"github.com/go-redis/redis/v9"
)

type redisCache struct {
	defaultExpiration time.Duration
	redisClient       *cache.RedisClient
}

func NewRedisCache(defaultExpiration time.Duration, redisClient *cache.RedisClient) *redisCache {
	return &redisCache{
		defaultExpiration: defaultExpiration,
		redisClient:       redisClient,
	}
}

func (c *redisCache) parseValue(value string) (cache.Item, error) {
	var item cache.Item
	err := json.Unmarshal([]byte(value), &item)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (c *redisCache) Get(key string) (interface{}, error) {
	val, err := c.redisClient.Get(key)
	if err != nil {
		if err == redis.Nil {
			return nil, common.ErrNoCacheKeyFound
		}
		return nil, err
	}

	item, err := c.parseValue(val)
	if err != nil {
		return nil, err
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return nil, common.ErrCacheKeyExpired
		}
	}

	return item.Object, nil
}

func (c *redisCache) Set(key string, value interface{}, d time.Duration) error {
	var e int64
	if d == cache.DefaultExpiration {
		d = c.defaultExpiration
	}

	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	item := cache.Item{
		Object:     value,
		Expiration: e,
	}
	byteValue, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return c.redisClient.Set(key, string(byteValue))
}

func (c *redisCache) Delete(key string) error {
	return c.redisClient.Delete(key)
}

func (c *redisCache) ItemCount() int {
	keys, err := c.redisClient.Keys("*")
	if err != nil {
		return 0
	}
	return len(keys)
}

func (c *redisCache) Items() map[string]cache.Item {
	keys, err := c.redisClient.Keys("*")
	if err != nil {
		return nil
	}
	items := make(map[string]cache.Item)
	temps := make(map[string]string)
	pipeliner := c.redisClient.Pipeline()

	for _, key := range keys {
		temps[key], err = pipeliner.Get(context.Background(), key).Result()
		if err != nil {
			return nil
		}
	}
	_, err = pipeliner.Exec(context.Background())
	if err != nil {
		return nil
	}
	for key, val := range temps {
		item, err := c.parseValue(val)
		if err != nil {
			continue
		}
		if item.Expiration > 0 {
			if time.Now().UnixNano() > item.Expiration {
				continue
			}
		}
		items[key] = item
	}
	return items
}
