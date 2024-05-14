package usecase

import "github.com/didiegovieira/go-benchmark-api/internal/domain/entity"

type PostSortingAlgorithmUseCaseInterface interface {
	Execute(arr []int) (*entity.Benchmark, error)
}
