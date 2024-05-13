package sortingalgorithm

type InsertionSortUseCase struct{}

func NewInsertionSortUseCase() *InsertionSortUseCase {
	return &InsertionSortUseCase{}
}

func (i *InsertionSortUseCase) Execute(arr []int) []int {
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
