package sortingalgorithm

import "github.com/didiegovieira/go-benchmark-api/pkg/base"

type InsertionSort = base.UseCase[[]int, []int]

type InsertionSortImplementation struct{}

func NewInsertionSort() *InsertionSortImplementation {
	return &InsertionSortImplementation{}
}

func (i *InsertionSortImplementation) Execute(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	return arr
}
