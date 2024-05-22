package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/didiegovieira/go-benchmark-api/internal/application/dto"
	"github.com/didiegovieira/go-benchmark-api/internal/domain/entity"
	"github.com/didiegovieira/go-benchmark-api/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPostSortingAlgorithmImplementationExecute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepositoryBenchmarkRepository(ctrl)
	mockTime := mocks.NewMockUseCase[dto.TimeCalculateInput, entity.Result](ctrl)

	sort := NewPostSortingAlgorithm(mockRepo, mockTime)

	input := dto.SortingInput{
		Arr: []int{3, 2, 1},
	}

	mockResult := &entity.Result{Name: "BubbleSort", Duration: time.Hour.Milliseconds()}

	mockTime.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(*mockResult, nil).AnyTimes()
	mockRepo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()

	result, err := sort.Execute(context.Background(), input)

	assert.NoError(t, err)

	assert.NotNil(t, result)
}
