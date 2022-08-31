package cache

import "time"

const (
	DefaultExpiration = 0
	NoExpiration      = -1
)

type Item struct {
	Object     interface{}
	Expiration int64
}

type ICache interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, d time.Duration) error
	Delete(key ...string) error
	DeleteByPattern(pattern string) error
	ItemCount() int
	Items() map[string]Item
}
