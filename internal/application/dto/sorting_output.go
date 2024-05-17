package dto

import "github.com/didiegovieira/go-benchmark-api/internal/domain/entity"

type SortingOutput struct {
	Message string           `json:"message"`
	Result  Result           `json:"result"`
	Data    entity.Benchmark `json:"data"`
}

type Result struct {
	Fastest string `json:"fastest"`
	Slowest string `json:"slowest"`
}

func (s *SortingOutput) NewSortingOutput(b entity.Benchmark) SortingOutput {
	return SortingOutput{
		Message: "Sorting Algorithm",
		Result: Result{
			Fastest: b.Faster.Duration.String(),
			Slowest: b.Slower.Duration.String(),
		},
		Data: b,
	}
}
