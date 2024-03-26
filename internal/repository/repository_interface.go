package repository

import (
	"context"

	"github.com/ali-shokoohi/notes/internal/model"
)

type CommonBehaviourRepository[T model.DBModel] interface {
	GetByFilterWithPagination(ctx context.Context, modelFilterData T, offset, limit int) ([]*T, error)
	Save(ctx context.Context, model *T) error
	Delete(ctx context.Context, model *T) error
}
