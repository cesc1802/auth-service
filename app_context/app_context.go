package app_context

import (
	"github.com/cesc1802/auth-service/pkg/logger"
	"github.com/cesc1802/auth-service/pkg/tokenprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetAppGorm() *gorm.DB
	GetATProvider() tokenprovider.Provider
	GetRTProvider() tokenprovider.Provider
	GetLogger() *logger.Logger
}

type appContext struct {
	db           *gorm.DB
	atProvider   tokenprovider.Provider
	rtProvider   tokenprovider.Provider
	customLogger *logger.Logger
}

func NewAppContext(db *gorm.DB, atProvider,
	rtProvider tokenprovider.Provider) *appContext {
	return &appContext{db: db, atProvider: atProvider, rtProvider: rtProvider, customLogger: logger.New("app_context")}
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
func (a *appContext) GetLogger() *logger.Logger {
	return a.customLogger
}
