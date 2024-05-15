package usecase

import (
	"context"
	"sort"

	"github.com/didiegovieira/go-benchmark-api/internal/application/dto"
	"github.com/didiegovieira/go-benchmark-api/internal/application/repository"
	sortingalgorithm "github.com/didiegovieira/go-benchmark-api/internal/application/usecase/sorting_algorithm"
	"github.com/didiegovieira/go-benchmark-api/internal/domain/entity"
	"github.com/didiegovieira/go-benchmark-api/pkg/base"
)

type SortingAlgorithm string

const (
	BubbleSort    SortingAlgorithm = "BubbleSort"
	InsertionSort SortingAlgorithm = "InsertionSort"
	MergeSort     SortingAlgorithm = "MergeSort"
	QuickSort     SortingAlgorithm = "QuickSort"
	SelectionSort SortingAlgorithm = "SelectionSort"
)

type PostSortingAlgorithmUseCase = base.UseCase[dto.RequestInput, *entity.Benchmark]

type PostSortingAlgorithmImplementation struct {
	repository           repository.BenchmarkRepository
	timeCalculateUseCase TimeCalculate
	sortingAlgorithms    map[SortingAlgorithm]sortingalgorithm.SortingAlgorithmsInterface
}

func NewPostSortingAlgorithm(
	repository repository.BenchmarkRepository,
	timeCalculateUseCase TimeCalculate,
	sortingAlgorithms map[SortingAlgorithm]sortingalgorithm.SortingAlgorithmsInterface,
) *PostSortingAlgorithmImplementation {
	return &PostSortingAlgorithmImplementation{
		repository:           repository,
		timeCalculateUseCase: timeCalculateUseCase,
		sortingAlgorithms: map[SortingAlgorithm]sortingalgorithm.SortingAlgorithmsInterface{
			BubbleSort:    sortingalgorithm.NewBubbleSortImplementation(),
			InsertionSort: sortingalgorithm.NewInsertionSortImplementation(),
			MergeSort:     sortingalgorithm.NewMergeSortImplementation(),
			QuickSort:     sortingalgorithm.NewQuickSortImplementation(),
			SelectionSort: sortingalgorithm.NewSelectionSortImplementation(),
		},
	}
}

func (s *PostSortingAlgorithmImplementation) Execute(ctx context.Context, input dto.RequestInput) (*entity.Benchmark, error) {
	b := s.initBenchmark(entity.SortingAlgorithm, input.Arr)
	s.calculateExecutionTimes(ctx, b, input.Arr)
	s.findFastestAndSlowest(b)
	err := s.saveToDatabase(b)

	return b, err
}

func (s *PostSortingAlgorithmImplementation) initBenchmark(bn entity.BenchmarkName, arr []int) *entity.Benchmark {
	b := &entity.Benchmark{}
	b.NewBenchmark(bn, arr)

	return b
}

func (s *PostSortingAlgorithmImplementation) saveToDatabase(b *entity.Benchmark) error {
	err := s.repository.Save(b)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostSortingAlgorithmImplementation) calculateExecutionTimes(ctx context.Context, b *entity.Benchmark, arr []int) {
	for algo, fn := range s.sortingAlgorithms {
		input := dto.TimeCalculateInput{
			Name: string(algo),
			Func: func() {
				_ = fn.Execute(arr)
			},
		}

		result, _ := s.timeCalculateUseCase.Execute(ctx, input)

		b.AddResult(result)
	}
}

func (s *PostSortingAlgorithmImplementation) findFastestAndSlowest(b *entity.Benchmark) {
	results := append([]entity.Result(nil), b.Results...)

	sort.Slice(results, func(i, j int) bool {
		return results[i].Duration < results[j].Duration
	})

	b.SetFast(results[0])
	b.SetSlow(results[len(results)-1])
}
