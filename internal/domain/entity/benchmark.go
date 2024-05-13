package entity

import (
	"time"

	"github.com/google/uuid"
)

type Type string

const (
	SortingAlgorithm Type = "sorting_algorithm"
	Serialization    Type = "serialization"
)

type Benchmark struct {
	Id      string    `json:"id"`
	Type    Type      `json:"type"`
	Data    []int     `json:"data"`
	Results []Result  `json:"results"`
	Fast    Result    `json:"fast"`
	Slow    Result    `json:"slow"`
	Date    time.Time `json:"date"`
}

type Result struct {
	Name     string        `json:"name"`
	Duration time.Duration `json:"duration"`
}

func (b *Benchmark) NewBenchmark(t Type, data []int) *Benchmark {
	return &Benchmark{
		Id:      uuid.New().String(),
		Type:    t,
		Data:    data,
		Results: []Result{},
		Date:    time.Now(),
	}
}
