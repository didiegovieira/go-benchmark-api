package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/didiegovieira/go-benchmark-api/internal/application/dto"
	"github.com/didiegovieira/go-benchmark-api/internal/domain/entity"
	cer "github.com/didiegovieira/go-benchmark-api/pkg/error"
	"github.com/didiegovieira/go-benchmark-api/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

func TestPostSortingAlgorithmHandleInvalidJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUseCase[dto.SortingInput, *entity.Benchmark](ctrl)
	mockPresenter := mocks.NewMockApiPresenter(ctrl)

	handler := &PostSortingAlgorithm{
		PostSortingAlgorithmUseCase: mockUseCase,
		Presenter:                   mockPresenter,
	}

	mockPresenter.EXPECT().Error(gomock.Any(), gomock.Any()).DoAndReturn(func(c *gin.Context, err *cer.Http) {
		c.JSON(err.Code, gin.H{"error": err.Message})
	}).Times(1)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/sort", handler.Handle())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/sort", bytes.NewBufferString(`invalid-json`))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPostSortingAlgorithmHandleUseCaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUseCase[dto.SortingInput, *entity.Benchmark](ctrl)
	mockPresenter := mocks.NewMockApiPresenter(ctrl)

	handler := &PostSortingAlgorithm{
		PostSortingAlgorithmUseCase: mockUseCase,
		Presenter:                   mockPresenter,
	}

	input := dto.SortingInput{ /* preencha com dados válidos */ }
	errCase := errors.New("use case error")
	mockUseCase.EXPECT().Execute(gomock.Any(), input).Return(nil, errCase).Times(1)
	mockPresenter.EXPECT().Error(gomock.Any(), gomock.Any()).DoAndReturn(func(c *gin.Context, err *cer.Http) {
		c.JSON(err.Code, gin.H{"error": err.Message})
	}).Times(1)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/sort", handler.Handle())

	jsonInput, _ := json.Marshal(input)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/sort", bytes.NewBuffer(jsonInput))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestPostSortingAlgorithmHandleSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mocks.NewMockUseCase[dto.SortingInput, *entity.Benchmark](ctrl)
	mockPresenter := mocks.NewMockApiPresenter(ctrl)

	handler := &PostSortingAlgorithm{
		PostSortingAlgorithmUseCase: mockUseCase,
		Presenter:                   mockPresenter,
	}

	input := dto.SortingInput{ /* preencha com dados válidos */ }
	benchmark := entity.Benchmark{ /* preencha com dados válidos */ }

	mockUseCase.EXPECT().Execute(gomock.Any(), input).Return(&benchmark, nil).Times(1)
	mockPresenter.EXPECT().Present(gomock.Any(), gomock.Any(), http.StatusOK).Times(1)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/sort", handler.Handle())

	jsonInput, _ := json.Marshal(input)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/sort", bytes.NewBuffer(jsonInput))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
