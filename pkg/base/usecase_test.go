package base_test

import (
	"context"
	"errors"
	"testing"

	"github.com/didiegovieira/go-benchmark-api/pkg/base"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type implUsecase struct {
}

type Usecase = base.UseCase[any, any]

func NewUsecase() Usecase {
	return new(implUsecase)
}

func (e implUsecase) Execute(ctx context.Context, _ any) (any, error) {
	return nil, nil
}

func TestUsingBaseUsecase(t *testing.T) {
	t.Run("Simulate UseCase Implementation with Generics", func(t *testing.T) {
		uc := NewUsecase()
		assert.NotNil(t, uc)

		output, err := uc.Execute(context.TODO(), nil)
		assert.Nil(t, output)
		assert.NoError(t, err)
	})
}

func TestUsingMockBaseUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUsecase := base.NewMockUseCase[implUsecase, any](ctrl)
	mockUsecase.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(1, nil)
	var uc = mockUsecase
	out, err := uc.Execute(context.TODO(), implUsecase{})
	assert.Equal(t, 1, out)
	assert.NoError(t, err)
}

func TestUsingMockBaseUsecaseReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUsecase := base.NewMockUseCase[implUsecase, any](ctrl)
	mockUsecase.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(nil, errors.New("something error"))
	var uc = mockUsecase
	out, err := uc.Execute(context.TODO(), implUsecase{})
	assert.Equal(t, nil, out)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "something")
}
