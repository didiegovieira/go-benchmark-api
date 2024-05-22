package repository

import "github.com/didiegovieira/go-benchmark-api/internal/domain/entity"

type BenchmarkRepository interface {
	Get(id string) (*entity.Benchmark, error)
	GetAll(benchmarkName string) ([]*entity.Benchmark, error)
	Save(benchmark *entity.Benchmark) error
}
