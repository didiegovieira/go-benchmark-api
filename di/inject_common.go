package di

import (
	"context"

	"github.com/didiegovieira/go-benchmark-api/internal/settings"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var commonSet = wire.NewSet()

func provideMongoDbClient() (*mongo.Client, func(), error) {
	mongodbClient, err := mongo.Connect(
		context.TODO(), options.Client().ApplyURI(settings.Settings.Database.Connection),
	)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		_ = mongodbClient.Disconnect(context.Background())
	}

	return mongodbClient, cleanup, nil
}
