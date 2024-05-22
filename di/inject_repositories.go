package di

import (
	repository2 "github.com/didiegovieira/go-benchmark-api/internal/application/repository"
	"github.com/didiegovieira/go-benchmark-api/internal/infrastructure/repository"
	"github.com/google/wire"
)

var provideBenchmarckRepositoryMongodb = wire.NewSet(
	repository.NewBenchmarkMongodb,
	wire.Bind(new(repository2.BenchmarkRepository), new(*repository.BenchmarkMongodb)),
)
