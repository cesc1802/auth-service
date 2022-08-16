package database

import (
	"fmt"
	"gorm.io/gorm"
)

func QueryPage(offset int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset - 1).Limit(limit)
	}
}

func Condition[T any](column string, value *T) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if value != nil {
			return db.Where(fmt.Sprintf("%s = ?", column), value)
		}
		return db
	}
}

func ActiveRecord(db *gorm.DB) *gorm.DB {
	return db.Where("status = ? and deleted_at != ?", 1, nil)
}
