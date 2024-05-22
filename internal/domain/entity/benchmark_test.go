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

func TestAddResult(t *testing.T) {
	benchmark := &Benchmark{}

	result := Result{Name: "Test Result", Duration: time.Hour.Milliseconds()}

	benchmark.AddResult(result)

	if len(benchmark.Results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(benchmark.Results))
	}
	if benchmark.Results[0] != result {
		t.Errorf("Expected result %+v, got %+v", result, benchmark.Results[0])
	}
}

func TestSetFast(t *testing.T) {
	benchmark := &Benchmark{}

	result := Result{Name: "Fast Result", Duration: time.Hour.Milliseconds()}

	benchmark.SetFast(result)

	if benchmark.Faster != result {
		t.Errorf("Expected fastest result %+v, got %+v", result, benchmark.Faster)
	}
}

func TestSetSlow(t *testing.T) {
	benchmark := &Benchmark{}

	result := Result{Name: "Slow Result", Duration: time.Hour.Milliseconds()}

	benchmark.SetSlow(result)

	if benchmark.Slower != result {
		t.Errorf("Expected slowest result %+v, got %+v", result, benchmark.Slower)
	}
}
