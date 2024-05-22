package base

import "context"

type Repository[M any] interface {
	FindOneById(ctx context.Context, id any) (M, error)
	InsertOne(ctx context.Context, model M) error
	DeleteOneById(ctx context.Context, id any) error
	UpInsert(ctx context.Context, id any, model M) error
}
