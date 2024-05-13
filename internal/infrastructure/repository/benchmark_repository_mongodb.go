package repository

import (
	"context"
	"time"

	"github.com/didiegovieira/go-benchmark-api/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BenchmarkModelMongodb struct {
	Id            string               `json:"id" bson:"_id"`
	BenchmarkName entity.BenchmarkName `json:"type" bson:"type"`
	Data          []int                `json:"data" bson:"data"`
	Results       []entity.Result      `json:"results" bson:"results"`
	Fast          entity.Result        `json:"fast" bson:"fast"`
	Slow          entity.Result        `json:"slow" bson:"slow"`
	Date          time.Time            `json:"date" bson:"date`
}

type BenchmarkRepositoryMongodb struct {
	Client     *mongo.Client
	collection *mongo.Collection
}

func NewBenchmarkRepositoryMongodb(client *mongo.Client, database string) *BenchmarkRepositoryMongodb {
	return &BenchmarkRepositoryMongodb{
		Client:     client,
		collection: client.Database(database).Collection("benchmark"),
	}
}

func (b *BenchmarkRepositoryMongodb) entityToModel(benchmarkEntity *entity.Benchmark) (benchmark BenchmarkModelMongodb) {
	return BenchmarkModelMongodb{
		Id:            benchmarkEntity.Id,
		BenchmarkName: benchmarkEntity.BenchmarkName,
		Data:          benchmarkEntity.Data,
		Results:       benchmarkEntity.Results,
		Fast:          benchmarkEntity.Fast,
		Slow:          benchmarkEntity.Slow,
		Date:          benchmarkEntity.Date,
	}
}

func (b *BenchmarkRepositoryMongodb) modelToEntity(benchmarkModel *BenchmarkModelMongodb) (benchmark *entity.Benchmark) {
	return &entity.Benchmark{
		Id:            benchmarkModel.Id,
		BenchmarkName: benchmarkModel.BenchmarkName,
		Data:          benchmarkModel.Data,
		Results:       benchmarkModel.Results,
		Fast:          benchmarkModel.Fast,
		Slow:          benchmarkModel.Slow,
		Date:          benchmarkModel.Date,
	}
}

func (o *BenchmarkRepositoryMongodb) Get(id string) (*entity.Benchmark, error) {
	var result BenchmarkModelMongodb

	if err := o.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return o.modelToEntity(&result), nil
}

func (o *BenchmarkRepositoryMongodb) GetAll(benchmarkName string) ([]*entity.Benchmark, error) {
	var results []*BenchmarkModelMongodb

	filter := bson.M{"benchmark_name": benchmarkName}

	cursor, err := o.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var result BenchmarkModelMongodb
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	benchmarks := make([]*entity.Benchmark, len(results))
	for i, bm := range results {
		benchmarks[i] = o.modelToEntity(bm)
	}

	return benchmarks, nil
}

func (b *BenchmarkRepositoryMongodb) Save(benchmark *entity.Benchmark) error {
	filter := bson.M{"_id": benchmark.Id}

	_, err := b.collection.ReplaceOne(
		context.Background(), filter, b.entityToModel(benchmark), options.Replace().SetUpsert(true),
	)

	if err != nil {
		return err
	}

	return nil
}
