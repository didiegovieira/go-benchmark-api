package sortingalgorithm

import (
	"reflect"
	"testing"
)

func TestMergeSortUseCaseExecute(t *testing.T) {
	// Create an instance of MergeSortUseCase
	mergeSortUseCase := NewMergeSort()

	// Test case 1: Unsorted array
	arr := []int{5, 3, 1, 4, 2}
	expected := []int{1, 2, 3, 4, 5}
	result := mergeSortUseCase.Execute(arr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect result. Expected: %v, Got: %v", expected, result)
	}

	// Test case 2: Already sorted array
	arr = []int{1, 2, 3, 4, 5}
	expected = []int{1, 2, 3, 4, 5}
	result = mergeSortUseCase.Execute(arr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect result. Expected: %v, Got: %v", expected, result)
	}

	// Test case 3: Reversed array
	arr = []int{5, 4, 3, 2, 1}
	expected = []int{1, 2, 3, 4, 5}
	result = mergeSortUseCase.Execute(arr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect result. Expected: %v, Got: %v", expected, result)
	}

	arr = []int{}
	expected = []int{}
	result = mergeSortUseCase.Execute(arr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect result. Expected: %v, Got: %v", expected, result)
	}
}
