package main

import (
	"github.com/didiegovieira/go-benchmark-api/internal/application/repository"
	usecase "github.com/didiegovieira/go-benchmark-api/internal/application/use_case"
	sortingalgorithm "github.com/didiegovieira/go-benchmark-api/internal/application/use_case/sorting_algorithm"
	"github.com/google/wire"
)

var useCasesSet = wire.NewSet(
	provideTimeCalculateUseCase,
	providePostSortingAlgorithmUseCase,
	provideBubbleSortUseCase,
	provideInsertionSortUseCase,
	provideMergeSortUseCase,
	provideQuickSortUseCase,
	provideSelectionSortUseCase,
)

func provideSelectionSortUseCase() sortingalgorithm.SelectionSortUseCaseInterface {
	return sortingalgorithm.NewSelectionSortUseCase()
}

func provideQuickSortUseCase() sortingalgorithm.QuickSortUseCaseInterface {
	return sortingalgorithm.NewQuickSortUseCase()
}

func provideMergeSortUseCase() sortingalgorithm.MergeSortUseCaseInterface {
	return sortingalgorithm.NewMergeSortUseCase()
}

func provideInsertionSortUseCase() sortingalgorithm.InsertionSortUseCaseInterface {
	return sortingalgorithm.NewInsertionSortUseCase()
}

func provideBubbleSortUseCase() sortingalgorithm.BubbleSortUseCaseInterface {
	return sortingalgorithm.NewBubbleSortUseCase()
}

func provideTimeCalculateUseCase() usecase.TimeCalculateUseCaseInterface {
	return usecase.NewTimeCalculateUseCase()
}

func providePostSortingAlgorithmUseCase(
	repository repository.BenchmarkRepositoryInterface,
	timeCalculateUseCase usecase.TimeCalculateUseCaseInterface,
	bubbleSortUseCase sortingalgorithm.BubbleSortUseCaseInterface,
	insertionSortUseCase sortingalgorithm.InsertionSortUseCaseInterface,
	mergeSortUseCase sortingalgorithm.MergeSortUseCaseInterface,
	quickSortUseCase sortingalgorithm.QuickSortUseCaseInterface,
	selectionSortUseCase sortingalgorithm.SelectionSortUseCaseInterface,
) usecase.PostSortingAlgorithmUseCaseInterface {
	return usecase.NewPostSortingAlgorithmUseCase(
		repository,
		timeCalculateUseCase,
		bubbleSortUseCase,
		insertionSortUseCase,
		mergeSortUseCase,
		quickSortUseCase,
		selectionSortUseCase,
	)
}
