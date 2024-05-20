package usecase

import (
	"context"
	"sort"

	"github.com/didiegovieira/go-benchmark-api/internal/application/dto"
	"github.com/didiegovieira/go-benchmark-api/internal/application/repository"
	"github.com/didiegovieira/go-benchmark-api/internal/application/usecase/sortingalgorithm"
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

type PostSortingAlgorithm = base.UseCase[dto.SortingInput, *entity.Benchmark]

type PostSortingAlgorithmImplementation struct {
	repository           repository.BenchmarkRepository
	timeCalculateUseCase TimeCalculate
}

func NewPostSortingAlgorithm(
	repository repository.BenchmarkRepository,
	timeCalculateUseCase TimeCalculate,
) *PostSortingAlgorithmImplementation {
	return &PostSortingAlgorithmImplementation{
		repository:           repository,
		timeCalculateUseCase: timeCalculateUseCase,
	}
}

func (s *PostSortingAlgorithmImplementation) Execute(ctx context.Context, input dto.SortingInput) (*entity.Benchmark, error) {
	b := s.initBenchmarkEntity(entity.SortingAlgorithm, input.Arr)
	s.calculateExecutionTimes(ctx, b, input.Arr)
	s.findFastestAndSlowest(b)
	err := s.saveToDatabase(b)

	return b, err
}

func (s *PostSortingAlgorithmImplementation) initBenchmarkEntity(bn entity.BenchmarkType, arr []int) *entity.Benchmark {
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
	for algo, fn := range s.mapAlgorithmFunctions() {
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

func (s *PostSortingAlgorithmImplementation) mapAlgorithmFunctions() map[SortingAlgorithm]sortingalgorithm.SortingAlgorithmsInterface {
	return map[SortingAlgorithm]sortingalgorithm.SortingAlgorithmsInterface{
		BubbleSort:    sortingalgorithm.NewBubbleSort(),
		InsertionSort: sortingalgorithm.NewInsertionSort(),
		MergeSort:     sortingalgorithm.NewMergeSort(),
		QuickSort:     sortingalgorithm.NewQuickSort(),
		SelectionSort: sortingalgorithm.NewSelectionSort(),
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
