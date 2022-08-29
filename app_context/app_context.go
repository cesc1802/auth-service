package app_context

import (
	"github.com/cesc1802/auth-service/pkg/broker"
	"github.com/cesc1802/auth-service/pkg/cache"
	"github.com/cesc1802/auth-service/pkg/tokenprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetAppGorm() *gorm.DB
	GetATProvider() tokenprovider.Provider
	GetRTProvider() tokenprovider.Provider
	GetAppCache() cache.ICache
	GetPublisher() broker.Publisher
	GetSubscriber() broker.Subscriber
}

type appContext struct {
	db         *gorm.DB
	atProvider tokenprovider.Provider
	rtProvider tokenprovider.Provider
	cache      cache.ICache
	publisher  broker.Publisher
	subscriber broker.Subscriber
}

func NewAppContext(db *gorm.DB, atProvider,
	rtProvider tokenprovider.Provider,
	cache cache.ICache,
	publisher broker.Publisher,
	subscriber broker.Subscriber,
) *appContext {
	return &appContext{
		db:         db,
		atProvider: atProvider,
		rtProvider: rtProvider,
		cache:      cache,
		publisher:  publisher,
		subscriber: subscriber,
	}
}

func (a *appContext) GetAppGorm() *gorm.DB {
	return a.db
}

func (a *appContext) GetATProvider() tokenprovider.Provider {
	return a.atProvider
}

func (a *appContext) GetRTProvider() tokenprovider.Provider {
	return a.rtProvider
}

func (a *appContext) GetAppCache() cache.ICache {
	return a.cache
}

func (a *appContext) GetPublisher() broker.Publisher {
	return a.publisher
}
func (a *appContext) GetSubscriber() broker.Subscriber {
	return a.subscriber
}
