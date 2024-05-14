package sortingalgorithm

type MergeSortUseCase struct{}

func NewMergeSortUseCase() *MergeSortUseCase {
	return &MergeSortUseCase{}
}

func (m *MergeSortUseCase) Execute(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	mid := n / 2
	left := m.Execute(arr[:mid])
	right := m.Execute(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	result = append(result, left...)
	result = append(result, right...)

	return result
}
