package base

import "context"

type UseCase[I any, O any] interface {
	Execute(ctx context.Context, input I) (O, error)
}
