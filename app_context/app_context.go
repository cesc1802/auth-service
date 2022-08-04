package app_context

import "gorm.io/gorm"

type AppContext interface {
	GetAppGorm() *gorm.DB
}

type appContext struct {
	db *gorm.DB
}

func NewAppContext() *appContext {
	return &appContext{}
}

func (a *appContext) GetAppGorm() *gorm.DB {
	return a.db
}
