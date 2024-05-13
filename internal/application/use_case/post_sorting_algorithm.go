package usecase

import (
	sortingalgorithm "github.com/didiegovieira/go-benchmark-api/internal/application/use_case/sorting_algorithm"
	"github.com/didiegovieira/go-benchmark-api/internal/domain/entity"
	"github.com/sirupsen/logrus"
)

type SortingAlgorithm string

const (
	BubbleSort    SortingAlgorithm = "BubbleSort"
	InsertionSort SortingAlgorithm = "InsertionSort"
	MergeSort     SortingAlgorithm = "MergeSort"
	QuickSort     SortingAlgorithm = "QuickSort"
	SelectionSort SortingAlgorithm = "SelectionSort"
)

type AlgorithmFunc func(arr []int) []int

type PostSortingAlgorithmUseCase struct {
	timeCalculateUseCase TimeCalculateUseCaseInterface
	sortingAlgorithms    map[SortingAlgorithm]AlgorithmFunc
}

func NewPostSortingAlgorithmUseCase(
	timeCalculateUseCase TimeCalculateUseCaseInterface,
	bubbleSortUseCase sortingalgorithm.BubbleSortUseCaseInterface,
	insertionSortUseCase sortingalgorithm.InsertionSortUseCaseInterface,
	mergeSortUseCase sortingalgorithm.MergeSortUseCaseInterface,
	quickSortUseCase sortingalgorithm.QuickSortUseCaseInterface,
	selectionSortUseCase sortingalgorithm.SelectionSortUseCaseInterface,
) *PostSortingAlgorithmUseCase {
	return &PostSortingAlgorithmUseCase{
		timeCalculateUseCase: timeCalculateUseCase,
		sortingAlgorithms: map[SortingAlgorithm]AlgorithmFunc{
			BubbleSort:    bubbleSortUseCase.Execute,
			InsertionSort: insertionSortUseCase.Execute,
			MergeSort:     mergeSortUseCase.Execute,
			QuickSort:     quickSortUseCase.Execute,
			SelectionSort: selectionSortUseCase.Execute,
		},
	}
}

func (s *PostSortingAlgorithmUseCase) Execute(arr []int) (b *entity.Benchmark) {
	b = b.NewBenchmark("sorting_algorithm", arr)
	s.timeCalculate(b, arr)
	s.fastAndSlow(b)

	logrus.Println("Fast: " + b.Fast.Name + " - " + b.Fast.Duration.String())
	logrus.Println("Slow: " + b.Slow.Name + " - " + b.Slow.Duration.String())

	return b
}

func (s *PostSortingAlgorithmUseCase) timeCalculate(b *entity.Benchmark, arr []int) {
	for algo, fn := range s.sortingAlgorithms {
		result := s.timeCalculateUseCase.Execute(func() {
			fn(arr)
		}, string(algo))
		b.Results = append(b.Results, result)
	}
}

func (s *PostSortingAlgorithmUseCase) fastAndSlow(b *entity.Benchmark) {
	fast := b.Results[0]
	slow := b.Results[0]

	for _, result := range b.Results {
		if result.Duration < fast.Duration {
			fast = result
		}

		if result.Duration > slow.Duration {
			slow = result
		}
	}

	b.Fast = fast
	b.Slow = slow
}
