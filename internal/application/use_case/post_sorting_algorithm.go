package usecase

import (
	"sort"

	"github.com/didiegovieira/go-benchmark-api/internal/application/repository"
	sortingalgorithm "github.com/didiegovieira/go-benchmark-api/internal/application/use_case/sorting_algorithm"
	"github.com/didiegovieira/go-benchmark-api/internal/domain/entity"
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
	repository           repository.BenchmarkRepositoryInterface
	timeCalculateUseCase TimeCalculateUseCaseInterface
	sortingAlgorithms    map[SortingAlgorithm]AlgorithmFunc
}

func NewPostSortingAlgorithmUseCase(
	repository repository.BenchmarkRepositoryInterface,
	timeCalculateUseCase TimeCalculateUseCaseInterface,
	bubbleSortUseCase sortingalgorithm.BubbleSortUseCaseInterface,
	insertionSortUseCase sortingalgorithm.InsertionSortUseCaseInterface,
	mergeSortUseCase sortingalgorithm.MergeSortUseCaseInterface,
	quickSortUseCase sortingalgorithm.QuickSortUseCaseInterface,
	selectionSortUseCase sortingalgorithm.SelectionSortUseCaseInterface,
) *PostSortingAlgorithmUseCase {
	return &PostSortingAlgorithmUseCase{
		repository:           repository,
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

func (s *PostSortingAlgorithmUseCase) Execute(arr []int) (*entity.Benchmark, error) {
	b := s.initBenchmark(entity.SortingAlgorithm, arr)

	s.calculateExecutionTimes(b, arr)

	s.findFastestAndSlowest(b)

	err := s.saveToDatabase(b)

	return b, err
}

func (s *PostSortingAlgorithmUseCase) saveToDatabase(b *entity.Benchmark) error {
	err := s.repository.Save(b)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostSortingAlgorithmUseCase) initBenchmark(bn entity.BenchmarkName, arr []int) *entity.Benchmark {
	b := &entity.Benchmark{}
	b.NewBenchmark(bn, arr)

	return b
}

func (s *PostSortingAlgorithmUseCase) calculateExecutionTimes(b *entity.Benchmark, arr []int) {
	for algo, fn := range s.sortingAlgorithms {
		result := s.timeCalculateUseCase.Execute(func() {
			fn(arr)
		}, string(algo))

		b.Results = append(b.Results, result)
	}
}

func (s *PostSortingAlgorithmUseCase) findFastestAndSlowest(b *entity.Benchmark) {
	results := append([]entity.Result(nil), b.Results...)

	sort.Slice(results, func(i, j int) bool {
		return results[i].Duration < results[j].Duration
	})

	b.Fast = results[0]
	b.Slow = results[len(results)-1]
}
