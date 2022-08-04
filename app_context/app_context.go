package app_context

import "gorm.io/gorm"

type AppContext interface {
	GetAppGorm() *gorm.DB
}

type appContext struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{db: db}
}

func (a *appContext) GetAppGorm() *gorm.DB {
	return a.db
}
