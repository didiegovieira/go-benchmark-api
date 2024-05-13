package usecase

import "github.com/didiegovieira/go-benchmark-api/internal/domain/entity"

type TimeCalculateUseCaseInterface interface {
	Execute(fn func(), name string) entity.Result
}
