package entity

import (
	"time"

	"github.com/google/uuid"
)

type BenchmarkName string

const (
	SortingAlgorithm BenchmarkName = "sorting_algorithm"
	Serialization    BenchmarkName = "serialization"
)

type Benchmark struct {
	Id            string        `json:"id"`
	BenchmarkName BenchmarkName `json:"benchmark_name"`
	Data          []int         `json:"data"`
	Results       []Result      `json:"results"`
	Fast          Result        `json:"fast"`
	Slow          Result        `json:"slow"`
	Date          time.Time     `json:"date"`
}

type Result struct {
	Name     string        `json:"name"`
	Duration time.Duration `json:"duration"`
}

func (b *Benchmark) NewBenchmark(bn BenchmarkName, data []int) {
	b.Id = uuid.New().String()
	b.BenchmarkName = bn
	b.Data = data
	b.Results = []Result{}
	b.Date = time.Now()
}

func (b *Benchmark) AddResult(result Result) {
	b.Results = append(b.Results, result)
}

func (b *Benchmark) SetFast(result Result) {
	b.Fast = result
}

func (b *Benchmark) SetSlow(result Result) {
	b.Slow = result
}
