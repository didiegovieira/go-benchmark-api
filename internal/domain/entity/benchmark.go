package entity

import (
	"time"

	"github.com/google/uuid"
)

type BenchmarkType string

const (
	SortingAlgorithm BenchmarkType = "sorting_algorithm"
	Serialization    BenchmarkType = "serialization"
)

type Benchmark struct {
	Id            string        `json:"id"`
	BenchmarkType BenchmarkType `json:"benchmark_type"`
	Data          []int         `json:"data"`
	Results       []Result      `json:"results"`
	Faster        Result        `json:"faster"`
	Slower        Result        `json:"slower"`
	CreatedAt     time.Time     `json:"created_at"`
}

type Result struct {
	Name     string        `json:"name"`
	Duration time.Duration `json:"duration"`
}

func (b *Benchmark) NewBenchmark(bn BenchmarkType, data []int) {
	b.Id = uuid.New().String()
	b.BenchmarkType = bn
	b.Data = data
	b.Results = []Result{}
	b.CreatedAt = time.Now()
}

func (b *Benchmark) AddResult(result Result) {
	b.Results = append(b.Results, result)
}

func (b *Benchmark) SetFast(result Result) {
	b.Faster = result
}

func (b *Benchmark) SetSlow(result Result) {
	b.Slower = result
}
