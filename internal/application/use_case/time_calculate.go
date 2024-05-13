package usecase

import (
	"time"

	"github.com/didiegovieira/go-benchmark-api/internal/domain/entity"
)

type TimeCalculateUseCase struct{}

func NewTimeCalculateUseCase() *TimeCalculateUseCase {
	return &TimeCalculateUseCase{}
}

func (s *TimeCalculateUseCase) Execute(fn func(), name string) entity.Result {
	start := time.Now()
	fn()

	resultTime := entity.Result{
		Name:     name,
		Duration: time.Since(start),
	}

	return resultTime
}
