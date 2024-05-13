package usecase

// import (
// 	"testing"
// 	"time"

// 	usecase "github.com/didiegovieira/go-benchmark-api/internal/application/use_case"
// 	"github.com/didiegovieira/go-benchmark-api/internal/domain/entity"
// 	mock_usecase "github.com/didiegovieira/go-benchmark-api/tests/mocks"
// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/mock/gomock"
// )

// func TestPostSortingAlgorithmUseCase_Execute(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockTimeCalculateUseCase := mock_usecase.NewMockTimeCalculateUseCaseInterface(ctrl)
// 	mockBubbleSortUseCase := mock_usecase.NewMockBubbleSortUseCaseInterface(ctrl)
// 	mockInsertionSortUseCase := mock_usecase.NewMockInsertionSortUseCaseInterface(ctrl)
// 	mockMergeSortUseCase := mock_usecase.NewMockMergeSortUseCaseInterface(ctrl)
// 	mockQuickSortUseCase := mock_usecase.NewMockQuickSortUseCaseInterface(ctrl)
// 	mockSelectionSortUseCase := mock_usecase.NewMockSelectionSortUseCaseInterface(ctrl)

// 	postSortingAlgorithmUseCase := usecase.NewPostSortingAlgorithmUseCase(
// 		mockTimeCalculateUseCase,
// 		mockBubbleSortUseCase,
// 		mockInsertionSortUseCase,
// 		mockMergeSortUseCase,
// 		mockQuickSortUseCase,
// 		mockSelectionSortUseCase,
// 	)

// 	arr := []int{5, 3, 8, 2, 1, 9, 4}

// 	mockTimeCalculateUseCase.EXPECT().Execute(gomock.Any(), gomock.Any()).Times(5).Return(entity.Result{
// 		Name:     "test",
// 		Duration: time.Second,
// 	}).Times(5)

// 	benchmark := postSortingAlgorithmUseCase.Execute(arr)

// 	assert.NotNil(t, benchmark)
// 	assert.Equal(t, 5, len(benchmark.Results))
//  }
