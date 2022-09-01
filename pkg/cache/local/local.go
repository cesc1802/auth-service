package local

import (
	"sync"
	"time"

	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/pkg/cache"
)

type localCache struct {
	defaultExpiration time.Duration
	items             map[string]cache.Item
	mu                sync.RWMutex
}

func NewLocalCache(defaultExpiration time.Duration) *localCache {
	return &localCache{
		defaultExpiration: defaultExpiration,
		items:             make(map[string]cache.Item),
	}
}

func (c *localCache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	item, found := c.items[key]
	if !found {
		c.mu.RUnlock()
		return nil, common.ErrNoCacheKeyFound
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return nil, common.ErrCacheKeyExpired
		}
	}
	c.mu.RUnlock()
	return item.Object, nil
}

func (c *localCache) Set(key string, value interface{}, d time.Duration) error {
	var e int64
	if d == cache.DefaultExpiration {
		d = c.defaultExpiration
	}

	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mu.Lock()
	c.items[key] = cache.Item{
		Object:     value,
		Expiration: e,
	}
	c.mu.Unlock()
	return nil
}

func (c *localCache) Delete(key string) error {
	c.mu.Lock()
	delete(c.items, key)
	c.mu.Unlock()
	return nil
}

func (c *localCache) DeleteByPattern(pattern string) error {
	return nil
}

func (c *localCache) ItemCount() int {
	c.mu.Lock()
	n := len(c.items)
	c.mu.Unlock()
	return n
}

func (c *localCache) Items() map[string]cache.Item {
	m := make(map[string]cache.Item)
	c.mu.Lock()
	for k, v := range c.items {
		if v.Expiration > 0 {
			if time.Now().UnixNano() > v.Expiration {
				continue
			}
		}
		m[k] = v
	}
	return m
}
