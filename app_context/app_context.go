package app_context

import (
	"github.com/cesc1802/auth-service/pkg/tokenprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetAppGorm() *gorm.DB
	GetATProvider() tokenprovider.Provider
	GetRTProvider() tokenprovider.Provider
}

type appContext struct {
	db         *gorm.DB
	atProvider tokenprovider.Provider
	rtProvider tokenprovider.Provider
}

func NewAppContext(db *gorm.DB, atProvider,
	rtProvider tokenprovider.Provider) *appContext {
	return &appContext{db: db, atProvider: atProvider, rtProvider: rtProvider}
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
