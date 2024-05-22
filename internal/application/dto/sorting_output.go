package dto

import "github.com/didiegovieira/go-benchmark-api/internal/domain/entity"

type SortingOutput struct {
	Message string           `json:"message"`
	Data    entity.Benchmark `json:"data"`
}

func (s *SortingOutput) NewSortingOutput(b entity.Benchmark) SortingOutput {
	return SortingOutput{
		Message: "Sorting Algorithm",
		Data:    b,
	}
}
