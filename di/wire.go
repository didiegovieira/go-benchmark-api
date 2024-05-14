//go:build wireinject
// +build wireinject

package main

import (
	usecase "github.com/didiegovieira/go-benchmark-api/internal/application/use_case"
	sortingalgorithm "github.com/didiegovieira/go-benchmark-api/internal/application/use_case/sorting_algorithm"
	"github.com/didiegovieira/go-benchmark-api/pkg/config"
	"github.com/didiegovieira/go-benchmark-api/pkg/route"
	"github.com/google/wire"
)

type DependencyContainer struct {
	Configs *config.Conf

	Routes []route.RouteInterface

	PostSortingAlgorithmUseCase usecase.PostSortingAlgorithmUseCaseInterface
	TimeCalculateUseCase        usecase.TimeCalculateUseCaseInterface
	BubbleSortUseCase           sortingalgorithm.BubbleSortUseCaseInterface
	InsertionSortUseCase        sortingalgorithm.InsertionSortUseCaseInterface
	MergeSortUseCase            sortingalgorithm.MergeSortUseCaseInterface
	QuickSortUseCase            sortingalgorithm.QuickSortUseCaseInterface
	SelectionSortUseCase        sortingalgorithm.SelectionSortUseCaseInterface
}

func newDependencyContainer(
	configs *config.Conf,

	routes []route.RouteInterface,

	postSortingAlgorithmUseCase usecase.PostSortingAlgorithmUseCaseInterface,
	timeCalculateUseCase usecase.TimeCalculateUseCaseInterface,
	bubbleSortUseCase sortingalgorithm.BubbleSortUseCaseInterface,
	insertionSortUseCase sortingalgorithm.InsertionSortUseCaseInterface,
	mergeSortUseCase sortingalgorithm.MergeSortUseCaseInterface,
	quickSortUseCase sortingalgorithm.QuickSortUseCaseInterface,
	selectionSortUseCase sortingalgorithm.SelectionSortUseCaseInterface,

) DependencyContainer {

	dc := DependencyContainer{
		Configs: configs,

		Routes: routes,
	}

	return dc
}

func InitializeDependencyContainer() (DependencyContainer, error) {
	wire.Build(
		config.LoadConfig,
		clientsSet,
		repositoriesSet,
		useCasesSet,
		routesSet,
		newDependencyContainer,
	)

	return DependencyContainer{}, nil
}
