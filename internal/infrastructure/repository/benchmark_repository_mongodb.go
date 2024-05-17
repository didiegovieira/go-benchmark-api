package repository

import (
	"context"
	"time"

	"github.com/didiegovieira/go-benchmark-api/internal/domain/entity"
	"github.com/didiegovieira/go-benchmark-api/internal/settings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BenchmarkModelMongodb struct {
	Id            string               `json:"id" bson:"_id"`
	BenchmarkType entity.BenchmarkType `json:"benchmark_type" bson:"benchmark_type"`
	Data          []int                `json:"data" bson:"data"`
	Results       []entity.Result      `json:"results" bson:"results"`
	Faster        entity.Result        `json:"faster" bson:"faster"`
	Slower        entity.Result        `json:"slower" bson:"slower"`
	CreatedAt     time.Time            `json:"created_at" bson:"created_at"`
}

type BenchmarkMongodb struct {
	Client     *mongo.Client
	collection *mongo.Collection
}

func NewBenchmarkMongodb(client *mongo.Client) *BenchmarkMongodb {
	benchmarkMongoDB := &BenchmarkMongodb{
		Client:     client,
		collection: client.Database(settings.Settings.Database.DbName).Collection("benchmark"),
	}

	_ = benchmarkMongoDB.createIndexes()

	return benchmarkMongoDB
}

func (b BenchmarkMongodb) createIndexes() error {
	_, err := b.collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{Keys: bson.D{{Key: "created_at", Value: 1}}},
	}, &options.CreateIndexesOptions{})

	if err != nil {
		return err
	}

	return nil
}

func (b *BenchmarkMongodb) entityToModel(benchmarkEntity *entity.Benchmark) (benchmark BenchmarkModelMongodb) {
	return BenchmarkModelMongodb{
		Id:            benchmarkEntity.Id,
		BenchmarkType: benchmarkEntity.BenchmarkType,
		Data:          benchmarkEntity.Data,
		Results:       benchmarkEntity.Results,
		Faster:        benchmarkEntity.Faster,
		Slower:        benchmarkEntity.Slower,
		CreatedAt:     benchmarkEntity.CreatedAt,
	}
}

func (b *BenchmarkMongodb) modelToEntity(benchmarkModel *BenchmarkModelMongodb) (benchmark *entity.Benchmark) {
	return &entity.Benchmark{
		Id:            benchmarkModel.Id,
		BenchmarkType: benchmarkModel.BenchmarkType,
		Data:          benchmarkModel.Data,
		Results:       benchmarkModel.Results,
		Faster:        benchmarkModel.Faster,
		Slower:        benchmarkModel.Slower,
		CreatedAt:     benchmarkModel.CreatedAt,
	}
}

func (o *BenchmarkMongodb) Get(id string) (*entity.Benchmark, error) {
	var result BenchmarkModelMongodb

	if err := o.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return o.modelToEntity(&result), nil
}

func (o *BenchmarkMongodb) GetAll(benchmarkName string) ([]*entity.Benchmark, error) {
	var results []*BenchmarkModelMongodb

	filter := bson.M{"benchmark_type": benchmarkName}

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

func (b *BenchmarkMongodb) Save(benchmark *entity.Benchmark) error {
	_, err := b.collection.InsertOne(context.Background(), b.entityToModel(benchmark))
	if err != nil {
		return err
	}

	return nil
}
