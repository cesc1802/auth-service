package generic

import (
	"context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/pkg/ifaces"
	"gorm.io/gorm"
)

type QueryFunc func(db *gorm.DB) *gorm.DB

type CRUDStore[T ifaces.Modeler] struct {
	db *gorm.DB
}

func (store *CRUDStore[TModel]) FindAll(ctx context.Context, queries ...QueryFunc) ([]TModel, error) {
	db := store.db

	var results []TModel

	if len(queries) > 0 {
		for _, handler := range queries {
			db = handler(db)
		}
	}

	var model *TModel
	if err := db.Model(model).Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (store *CRUDStore[TModel]) Update(ctx context.Context, model *TModel, queries ...QueryFunc) error {
	tx := store.db.Begin()

	if len(queries) > 0 {
		for _, handler := range queries {
			tx = handler(tx)
		}
	}

	if err := tx.Model(model).Updates(model).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(model).Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (store *CRUDStore[TModel]) Delete(ctx context.Context, id uint, queries ...QueryFunc) error {
	tx := store.db.Begin()
	var model TModel

	if len(queries) > 0 {
		for _, handler := range queries {
			tx = handler(tx)
		}
	}

	if err := tx.Model(model).Where("id = ?", id).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}

func (store *CRUDStore[TModel]) FindOne(ctx context.Context, id uint, queries ...QueryFunc) (*TModel, error) {
	db := store.db
	var result TModel

	if len(queries) > 0 {
		for _, handler := range queries {
			db = handler(db)
		}
	}

	if err := db.Model(result).Where("id = ?", id).Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
	}

	return &result, nil
}

func (store *CRUDStore[TModel]) FindOneByCondition(ctx context.Context, queries ...QueryFunc) (*TModel, error) {
	db := store.db
	var result TModel

	if len(queries) > 0 {
		for _, handler := range queries {
			db = handler(db)
		}
	}

	if err := db.Model(result).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
	}

	return &result, nil
}

func (store *CRUDStore[TModel]) Create(ctx context.Context, model *TModel, queries ...QueryFunc) error {
	tx := store.db.Begin()

	if len(queries) > 0 {
		for _, handler := range queries {
			tx = handler(tx)
		}
	}

	if err := tx.Model(model).Create(model).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(model).Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func NewCRUDStore[TModel ifaces.Modeler](db *gorm.DB) *CRUDStore[TModel] {
	return &CRUDStore[TModel]{db: db}
}
