package generic

import "context"

type IFindStore[T any] interface {
	FindOne(ctx context.Context, id uint, queries ...QueryFunc) (*T, error)
}

type IFindOneByConditionStore[T any] interface {
	FindOneByCondition(ctx context.Context, queries ...QueryFunc) (*T, error)
}

type ICreateStore[T any] interface {
	Create(ctx context.Context, model *T, queries ...QueryFunc) error
}

type IUpdateStore[T any] interface {
	Update(ctx context.Context, model *T, queries ...QueryFunc) error
}

type IFindAllStore[T any] interface {
	FindAll(ctx context.Context, queries ...QueryFunc) ([]T, *int64, error)
}

type IDeleteStore[T any] interface {
	Delete(ctx context.Context, id uint, queries ...QueryFunc) error
}

type ICRUDStore[T any] interface {
	IFindStore[T]
	ICreateStore[T]
	IFindAllStore[T]
	IUpdateStore[T]
	IDeleteStore[T]
}
