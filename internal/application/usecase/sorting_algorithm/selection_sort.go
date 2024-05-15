package sortingalgorithm

import "github.com/didiegovieira/go-benchmark-api/pkg/base"

type SelectionSort = base.UseCase[[]int, []int]

type SelectionSortImplementation struct{}

func NewSelectionSortImplementation() *SelectionSortImplementation {
	return &SelectionSortImplementation{}
}

func (s *SelectionSortImplementation) Execute(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}
