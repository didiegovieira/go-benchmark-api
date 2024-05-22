package api

import "context"

type Server[T any] interface {
	GetRouter() T
	Start() error
	Shutdown(ctx context.Context) error
}
