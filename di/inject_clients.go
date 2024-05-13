package main

import (
	"context"

	"github.com/didiegovieira/go-benchmark-api/pkg/config"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientsSet = wire.NewSet()

func provideMongoDbClient(configs *config.Conf) *mongo.Client {
	mongoUri := configs.GetConfig("MONGODB_URI", "")

	if mongoUri == "" {
		panic("MONGODB_URI is not set")
	}

	mongodbClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}

	return mongodbClient
}
