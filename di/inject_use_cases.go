package di

import (
	"github.com/didiegovieira/go-benchmark-api/internal/application/dto"
	"github.com/didiegovieira/go-benchmark-api/internal/application/repository"
	"github.com/didiegovieira/go-benchmark-api/internal/application/usecase"
	"github.com/didiegovieira/go-benchmark-api/internal/domain/entity"
	"github.com/didiegovieira/go-benchmark-api/pkg/base"
)

func providePostSortingAlgorithmUseCase(
	repository repository.BenchmarkRepository,
	timeCalculate usecase.TimeCalculate,
) base.UseCase[dto.SortingInput, *entity.Benchmark] {
	return usecase.NewPostSortingAlgorithm(
		repository,
		timeCalculate,
	)
}

func provideTimeCalculateUseCase() base.UseCase[dto.TimeCalculateInput, entity.Result] {
	return usecase.NewTimeCalculate()
}
