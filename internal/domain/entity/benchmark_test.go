package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var benchmark *Benchmark

func TestBenchmarkNewBenchmark(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	benchmarkType := SortingAlgorithm

	benchmark := benchmark.NewBenchmark(benchmarkType, data)

	assert.NotNil(t, benchmark)
	assert.NotEmpty(t, benchmark.Id)
	assert.Equal(t, benchmarkType, benchmark.BenchmarkName)
	assert.Equal(t, data, benchmark.Data)
	assert.Empty(t, benchmark.Results)
	assert.True(t, benchmark.Date.Before(time.Now()))
}

func TestBenchmarkNewBenchmarkSerialization(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	benchmarkType := Serialization

	benchmark := benchmark.NewBenchmark(benchmarkType, data)

	assert.NotNil(t, benchmark)
	assert.NotEmpty(t, benchmark.Id)
	assert.Equal(t, benchmarkType, benchmark.BenchmarkName)
	assert.Equal(t, data, benchmark.Data)
	assert.Empty(t, benchmark.Results)
	assert.True(t, benchmark.Date.Before(time.Now()))
}

func TestBenchmarkNewBenchmarkUniqueId(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	benchmarkType := SortingAlgorithm

	benchmark1 := benchmark.NewBenchmark(benchmarkType, data)
	benchmark2 := benchmark.NewBenchmark(benchmarkType, data)

	assert.NotEqual(t, benchmark1.Id, benchmark2.Id)
}

func TestBenchmarkNewBenchmarkDate(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	benchmarkType := SortingAlgorithm

	benchmark := benchmark.NewBenchmark(benchmarkType, data)

	assert.True(t, benchmark.Date.Before(time.Now()))
}
