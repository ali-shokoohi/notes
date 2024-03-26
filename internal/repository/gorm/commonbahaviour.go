package gorm

import (
	"context"

	"github.com/ali-shokoohi/notes/internal/errors"
	"github.com/ali-shokoohi/notes/internal/model"
	"github.com/ali-shokoohi/notes/internal/repository"
	"gorm.io/gorm"
)

type commonBehaviour[T model.DBModel] struct {
	db *gorm.DB
}

func NewCommonBehaviour[T model.DBModel](db *gorm.DB) repository.CommonBehaviourRepository[T] {
	return &commonBehaviour[T]{
		db: db,
	}
}

func (c *commonBehaviour[T]) GetByFilterWithPagination(ctx context.Context, modelFilterData T, offset, limit int) ([]*T, error) {
	var t []*T
	err := c.db.WithContext(ctx).Where(modelFilterData).Offset(offset).Limit(limit).Find(&t).Error
	if err != nil {
		// Convert Gorm error to custom application error
		return nil, errors.ConvertGormError(err)
	}
	return t, err
}

func (c *commonBehaviour[T]) Save(ctx context.Context, model *T) error {
	err := c.db.WithContext(ctx).Save(model).Error
	if err != nil {
		// Convert Gorm error to custom application error
		return errors.ConvertGormError(err)
	}
	return nil
}

func (c *commonBehaviour[T]) Delete(ctx context.Context, model *T) error {
	err := c.db.WithContext(ctx).Delete(model).Error
	if err != nil {
		// Convert Gorm error to custom application error
		return errors.ConvertGormError(err)
	}
	return nil
}
