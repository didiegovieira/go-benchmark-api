package repository

import "github.com/didiegovieira/go-benchmark-api/internal/domain/entity"

type BenchmarkRepositoryInterface interface {
	Save(benchmark *entity.Benchmark) error
}
