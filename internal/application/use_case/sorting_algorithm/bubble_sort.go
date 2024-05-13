package sortingalgorithm

type BubbleSortUseCase struct{}

func NewBubbleSortUseCase() *BubbleSortUseCase {
	return &BubbleSortUseCase{}
}

func (b *BubbleSortUseCase) Execute(arr []int) []int {
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
