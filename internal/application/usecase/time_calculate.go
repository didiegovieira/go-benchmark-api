package usecase

import (
	"context"
	"time"

	"github.com/didiegovieira/go-benchmark-api/internal/application/dto"
	"github.com/didiegovieira/go-benchmark-api/internal/domain/entity"
	"github.com/didiegovieira/go-benchmark-api/pkg/base"
)

type TimeCalculate = base.UseCase[dto.TimeCalculateInput, entity.Result]

type TimeCalculateImplementation struct{}

func NewTimeCalculateImplementation() *TimeCalculateImplementation {
	return &TimeCalculateImplementation{}
}

func (s *TimeCalculateImplementation) Execute(ctx context.Context, input dto.TimeCalculateInput) entity.Result {
	start := time.Now()
	input.Func()

	resultTime := s.creatingResultEntity(input.Name, time.Since(start))

	return resultTime
}

func (s *TimeCalculateImplementation) creatingResultEntity(name string, duration time.Duration) entity.Result {
	return entity.Result{
		Name:     name,
		Duration: duration,
	}
}
