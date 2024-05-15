package sortingalgorithm

import "github.com/didiegovieira/go-benchmark-api/pkg/base"

type BubbleSort = base.UseCase[[]int, []int]

type BubbleSortImplementation struct{}

func NewBubbleSortImplementation() *BubbleSortImplementation {
	return &BubbleSortImplementation{}
}

func (b *BubbleSortImplementation) Execute(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
