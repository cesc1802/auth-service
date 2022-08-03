package database

import (
	"fmt"
	"gorm.io/gorm"
)

//func WithCondition[T any](value T) func(db *gorm.DB) *gorm.DB {
//	return func(db *gorm.DB) *gorm.DB {
//		return db.Where("")
//	}
//}
//
//func Where(value string) *gorm.DB {
//	return WithCondition[string](value)()
//}

func WithCondition(conditions map[string]interface{}) func(query *gorm.DB) *gorm.DB {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(conditions)
	}
}

func QueryByLoginID(loginID string) func(query *gorm.DB) *gorm.DB {
	return WithCondition(map[string]interface{}{
		"login_id": loginID,
	})
}

func QueryByUserID(userID uint) func(query *gorm.DB) *gorm.DB {
	return WithCondition(map[string]interface{}{
		"id": userID,
	})
}

func QueryByPrimaryKey(ID uint) func(query *gorm.DB) *gorm.DB {
	return WithCondition(map[string]interface{}{
		"id": ID,
	})
}

func QueryPreload(values []string) func(query *gorm.DB) *gorm.DB {
	return func(query *gorm.DB) *gorm.DB {
		if len(values) > 0 {
			for _, value := range values {
				query = query.Preload(value)
			}
		}
		return query
	}

}

func QueryByStatus(status int) func(query *gorm.DB) *gorm.DB {
	return WithCondition(map[string]interface{}{
		"status": status,
	})
}
func IsNotAdmin() func(query *gorm.DB) *gorm.DB {
	return WithCondition(map[string]interface{}{
		"is_admin": 0,
	})
}

func QueryByVerifyLevel(level []int) func(query *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("verify_level IN ?", level)
	}
}

func QueryContaining(col string, value string) func(query *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s like ?", col), "%"+value+"%")
	}
}
