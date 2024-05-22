package repository

import (
	"os"
	"testing"
	"time"

	"github.com/didiegovieira/go-benchmark-api/internal/domain/entity"
	"github.com/didiegovieira/go-benchmark-api/internal/settings"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

var clientOpts = mtest.NewOptions().ClientType(mtest.Mock).DatabaseName(settings.Settings.Database.DbName)

func Init() {
	os.Setenv("MONGODB_URI", "mongodb://localhost:27017")
	os.Setenv("MONGODB_DB_NAME", "test")
	settings.Init()
}

func List() []*entity.Benchmark {
	return []*entity.Benchmark{
		{
			Id:            "123",
			BenchmarkType: entity.SortingAlgorithm,
			Data:          []int{1, 2, 3},
			Results:       []entity.Result{{Name: "Test", Duration: 10}},
			Faster:        entity.Result{Name: "Faster", Duration: 5},
			Slower:        entity.Result{Name: "Slower", Duration: 15},
			CreatedAt:     time.Now(),
		},
		{
			Id:            "1234",
			BenchmarkType: entity.SortingAlgorithm,
			Data:          []int{1, 2, 3},
			Results:       []entity.Result{{Name: "Test", Duration: 10}},
			Faster:        entity.Result{Name: "Faster", Duration: 5},
			Slower:        entity.Result{Name: "Slower", Duration: 15},
			CreatedAt:     time.Now(),
		},
	}
}

func TestNewBenchmarkMongodb(t *testing.T) {
	mt := mtest.New(t, clientOpts)

	mt.Run(t.Name(), func(mt *mtest.T) {
		repository := NewBenchmarkMongodb(mt.Client)
		assert.Equal(t, repository.Client, mt.Client)
		assert.Equal(t, mt.DB.Name(), "test")
		assert.Equal(t, repository.collection.Name(), "benchmark")
	})
}

func TestCreateIndexes(t *testing.T) {
	Init()

	mt := mtest.New(t, clientOpts)

	mt.Run(t.Name(), func(mt *mtest.T) {
		repository := NewBenchmarkMongodb(mt.Client)

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{}},
		})

		err := repository.createIndexes()

		assert.Nil(t, err)
	})
}

func TestEntityToModel(t *testing.T) {
	benchmarkEntity := &entity.Benchmark{
		Id:            "123",
		BenchmarkType: entity.SortingAlgorithm,
		Data:          []int{1, 2, 3},
		Results:       []entity.Result{{Name: "Test", Duration: 10}},
		Faster:        entity.Result{Name: "Faster", Duration: 5},
		Slower:        entity.Result{Name: "Slower", Duration: 15},
		CreatedAt:     time.Now(),
	}

	repo := &BenchmarkMongodb{}
	benchmarkModel := repo.entityToModel(benchmarkEntity)

	assert.NotNil(t, benchmarkModel)
	assert.Equal(t, "123", benchmarkModel.Id)
	assert.Equal(t, entity.SortingAlgorithm, benchmarkModel.BenchmarkType)
	assert.Equal(t, []int{1, 2, 3}, benchmarkModel.Data)
	assert.Len(t, benchmarkModel.Results, 1)
	assert.Equal(t, "Test", benchmarkModel.Results[0].Name)
	assert.Equal(t, int64(10), benchmarkModel.Results[0].Duration)
	assert.Equal(t, "Faster", benchmarkModel.Faster.Name)
	assert.Equal(t, int64(5), benchmarkModel.Faster.Duration)
	assert.Equal(t, "Slower", benchmarkModel.Slower.Name)
	assert.Equal(t, int64(15), benchmarkModel.Slower.Duration)
}

func TestModelToEntity(t *testing.T) {
	benchmarkModel := &BenchmarkModelMongodb{
		Id:            "123",
		BenchmarkType: entity.SortingAlgorithm,
		Data:          []int{1, 2, 3},
		Results:       []entity.Result{{Name: "Test", Duration: 10}},
		Faster:        entity.Result{Name: "Faster", Duration: 5},
		Slower:        entity.Result{Name: "Slower", Duration: 15},
		CreatedAt:     time.Now(),
	}

	repo := &BenchmarkMongodb{}
	benchmarkEntity := repo.modelToEntity(benchmarkModel)

	assert.NotNil(t, benchmarkEntity)
	assert.Equal(t, "123", benchmarkEntity.Id)
	assert.Equal(t, entity.SortingAlgorithm, benchmarkEntity.BenchmarkType)
	assert.Equal(t, []int{1, 2, 3}, benchmarkEntity.Data)
	assert.Len(t, benchmarkEntity.Results, 1)
	assert.Equal(t, "Test", benchmarkEntity.Results[0].Name)
	assert.Equal(t, int64(10), benchmarkEntity.Results[0].Duration)
	assert.Equal(t, "Faster", benchmarkEntity.Faster.Name)
	assert.Equal(t, int64(5), benchmarkEntity.Faster.Duration)
	assert.Equal(t, "Slower", benchmarkEntity.Slower.Name)
	assert.Equal(t, int64(15), benchmarkEntity.Slower.Duration)
}

func TestGet(t *testing.T) {
	Init()

	mt := mtest.New(t, clientOpts)

	mt.Run(t.Name(), func(mt *mtest.T) {
		repository := NewBenchmarkMongodb(mt.Client)

		expectedID := "123"
		expectedBenchmark := &BenchmarkModelMongodb{
			Id:            expectedID,
			BenchmarkType: entity.SortingAlgorithm,
			Data:          []int{1, 2, 3},
			Results:       []entity.Result{{Name: "Test", Duration: 10}},
			Faster:        entity.Result{Name: "Faster", Duration: 5},
			Slower:        entity.Result{Name: "Slower", Duration: 15},
			CreatedAt:     time.Now(),
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedBenchmark.Id},
			{Key: "benchmark_type", Value: expectedBenchmark.BenchmarkType},
			{Key: "data", Value: expectedBenchmark.Data},
			{Key: "results", Value: expectedBenchmark.Results},
			{Key: "faster", Value: expectedBenchmark.Faster},
			{Key: "slower", Value: expectedBenchmark.Slower},
			{Key: "created_at", Value: expectedBenchmark.CreatedAt},
		}))

		result, err := repository.Get(expectedID)

		assert.Nil(t, err)

		assert.NotNil(t, result)
		assert.Equal(t, expectedID, result.Id)
		assert.Equal(t, expectedBenchmark.BenchmarkType, result.BenchmarkType)
		assert.Equal(t, expectedBenchmark.Data, result.Data)
		assert.Equal(t, expectedBenchmark.Results, result.Results)
		assert.Equal(t, expectedBenchmark.Faster, result.Faster)
		assert.Equal(t, expectedBenchmark.Slower, result.Slower)
		assert.WithinDuration(t, expectedBenchmark.CreatedAt, result.CreatedAt, 1*time.Second)
	})
}

func TestGetNotFound(t *testing.T) {
	Init()

	mt := mtest.New(t, clientOpts)

	mt.Run(t.Name(), func(mt *mtest.T) {
		repository := NewBenchmarkMongodb(mt.Client)

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch))

		result, err := repository.Get("123")

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestGetErrMongo(t *testing.T) {
	Init()

	mt := mtest.New(t, clientOpts)

	mt.Run(t.Name(), func(mt *mtest.T) {
		repository := NewBenchmarkMongodb(mt.Client)

		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "mongodb error",
		}))

		result, err := repository.Get("123")

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, err.Error(), "mongodb error")
	})
}

func TestGetAll(t *testing.T) {
	Init()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("TestGetAll", func(mt *mtest.T) {
		repository := NewBenchmarkMongodb(mt.Client)

		expectedBenchmarks := List()

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedBenchmarks[0].Id},
			{Key: "benchmark_type", Value: expectedBenchmarks[0].BenchmarkType},
			{Key: "data", Value: expectedBenchmarks[0].Data},
			{Key: "results", Value: expectedBenchmarks[0].Results},
			{Key: "faster", Value: expectedBenchmarks[0].Faster},
			{Key: "slower", Value: expectedBenchmarks[0].Slower},
			{Key: "created_at", Value: expectedBenchmarks[0].CreatedAt},
		}))

		results, err := repository.GetAll("sorting_algorithm")

		assert.Nil(t, err)
		assert.NotNil(t, results)
		assert.Equal(t, expectedBenchmarks[0].Id, results[0].Id)
	})
}

func TestGetAllFindError(t *testing.T) {
	Init()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("TestGetAllFindError", func(mt *mtest.T) {
		repository := NewBenchmarkMongodb(mt.Client)

		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    1,
			Message: "mock find error",
		}))

		results, err := repository.GetAll("sorting_algorithm")

		assert.NotNil(t, err)
		assert.Nil(t, results)
		assert.Contains(t, err.Error(), "mock find error")
	})
}

func TestGetAllErrMongo(t *testing.T) {
	Init()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("TestGetAll", func(mt *mtest.T) {
		repository := NewBenchmarkMongodb(mt.Client)

		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "mongodb error",
		}))

		results, err := repository.GetAll("sorting_algorithm")

		assert.NotNil(t, err)
		assert.Nil(t, results)
		assert.Equal(t, err.Error(), "mongodb error")
	})
}

func TestSave(t *testing.T) {
	Init()

	mt := mtest.New(t, clientOpts)

	mt.Run(t.Name(), func(mt *mtest.T) {
		repository := NewBenchmarkMongodb(mt.Client)

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{}},
		})

		err := repository.Save(&entity.Benchmark{})

		assert.Nil(t, err)
	})
}

func TestSaveWithError(t *testing.T) {
	Init()

	mt := mtest.New(t, clientOpts)

	mt.Run(t.Name(), func(mt *mtest.T) {
		repository := NewBenchmarkMongodb(mt.Client)

		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "mongodb error",
		}))

		err := repository.Save(&entity.Benchmark{})

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "mongodb error")
	})
}
