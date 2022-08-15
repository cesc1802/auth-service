package app_context

import (
	"github.com/cesc1802/auth-service/pkg/cache"
	"gorm.io/gorm"
)

type AppContext interface {
	GetAppGorm() *gorm.DB
	GetAppCache() cache.ICache
}

type appContext struct {
	db    *gorm.DB
	cache cache.ICache
}

func NewAppContext(db *gorm.DB, cache cache.ICache) *appContext {
	return &appContext{db: db, cache: cache}
}

func (a *appContext) GetAppGorm() *gorm.DB {
	return a.db
}

func (a *appContext) GetAppCache() cache.ICache {
	return a.cache
}
