package usecase

import (
	"testing"

	"github.com/didiegovieira/go-benchmark-api/tests/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	mock "go.uber.org/mock/gomock"
)

func TestTimeCalculateUseCaseExecute(t *testing.T) {
	// Create an instance of TimeCalculateUseCase
	timeCalculateUseCase := NewTimeCalculateUseCase()
	arr := []int{1, 2, 3}

	// Define a function to measure the execution time
	ctrl := mock.NewController(t)
	testFunc := mocks.NewMockBubbleSortUseCaseInterface(ctrl)
	testFunc.EXPECT().Execute(arr).Return(arr)

	// Execute the function and measure the time
	result := timeCalculateUseCase.Execute(func() { testFunc.Execute(arr) }, "BubbleSort")

	assert.NotEmpty(t, result.Name)
	logrus.Println("Name: " + result.Name)
	assert.NotEmpty(t, result.Duration)
	logrus.Println("Duration: " + result.Duration.String())
}
