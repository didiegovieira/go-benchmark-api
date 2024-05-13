package sortingalgorithm

type QuickSortUseCase struct{}

func NewQuickSortUseCase() *QuickSortUseCase {
	return &QuickSortUseCase{}
}

func (q *QuickSortUseCase) Execute(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1]
	var left, right []int
	for _, v := range arr[:len(arr)-1] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = q.Execute(left)
	right = q.Execute(right)

	return append(append(left, pivot), right...)
}
