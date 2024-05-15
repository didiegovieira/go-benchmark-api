package main

import (
	"github.com/didiegovieira/go-benchmark-api/internal/application/repository"
	usecase "github.com/didiegovieira/go-benchmark-api/internal/application/use_case"
	"github.com/google/wire"
)

var useCasesSet = wire.NewSet(
	provideTimeCalculateUseCase,
	providePostSortingAlgorithmUseCase,
)

func provideTimeCalculateUseCase() usecase.TimeCalculateUseCaseInterface {
	return usecase.NewTimeCalculateUseCase()
}

func providePostSortingAlgorithmUseCase(
	repository repository.BenchmarkRepositoryInterface,
	timeCalculateUseCase usecase.TimeCalculateUseCaseInterface,

) usecase.PostSortingAlgorithmUseCaseInterface {
	return usecase.NewPostSortingAlgorithmUseCase(
		repository,
		timeCalculateUseCase,
	)
}
