package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBenchmarkNewBenchmark(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	benchmarkType := SortingAlgorithm

	benchmark := Benchmark{}
	benchmark.NewBenchmark(benchmarkType, data)

	assert.NotNil(t, benchmark)
	assert.NotEmpty(t, benchmark.Id)
	assert.Equal(t, benchmarkType, benchmark.BenchmarkType)
	assert.Equal(t, data, benchmark.Data)
	assert.Empty(t, benchmark.Results)
	assert.True(t, benchmark.CreatedAt.Before(time.Now()))
}

func TestBenchmarkNewBenchmarkSerialization(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	benchmarkType := Serialization

	benchmark := Benchmark{}
	benchmark.NewBenchmark(benchmarkType, data)

	assert.NotNil(t, benchmark)
	assert.NotEmpty(t, benchmark.Id)
	assert.Equal(t, benchmarkType, benchmark.BenchmarkType)
	assert.Equal(t, data, benchmark.Data)
	assert.Empty(t, benchmark.Results)
	assert.True(t, benchmark.CreatedAt.Before(time.Now()))
}

func TestBenchmarkNewBenchmarkUniqueId(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	benchmarkType := SortingAlgorithm

	benchmark1 := Benchmark{}
	benchmark2 := Benchmark{}

	benchmark1.NewBenchmark(benchmarkType, data)
	benchmark2.NewBenchmark(benchmarkType, data)

	assert.NotEqual(t, benchmark1.Id, benchmark2.Id)
}

func TestBenchmarkNewBenchmarkDate(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	benchmarkType := SortingAlgorithm

	benchmark := Benchmark{}
	benchmark.NewBenchmark(benchmarkType, data)

	assert.True(t, benchmark.CreatedAt.Before(time.Now()))
}
