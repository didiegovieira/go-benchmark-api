package repository

import (
	"github.com/didiegovieira/go-benchmark-api/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type BenchmarkModelMongodb struct {
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

func (b *BenchmarkRepositoryMongodb) entityToModel() (benchmark *BenchmarkModelMongodb) {
	return nil
}

func (b *BenchmarkRepositoryMongodb) modelToEntity() (benchmark *entity.Benchmark) {
	return nil
}

func (b *BenchmarkRepositoryMongodb) Save() error {
	return nil
}
