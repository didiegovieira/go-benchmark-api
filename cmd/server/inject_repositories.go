package main

import (
	"github.com/didiegovieira/go-benchmark-api/internal/application/repository"
	database "github.com/didiegovieira/go-benchmark-api/internal/infrastructure/repository"

	"github.com/didiegovieira/go-benchmark-api/pkg/config"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var repositoriesSet = wire.NewSet(
	provideBenchmarkRepository,
)

func provideBenchmarkRepository(client *mongo.Client, conf *config.Conf) repository.BenchmarkRepositoryInterface {
	return database.NewBenchmarkRepositoryMongodb(client, conf.GetOrPanic("MONGODB_DATABASE", "MONGODB_DATABASE is required"))
}
