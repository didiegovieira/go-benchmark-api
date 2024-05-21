//go:build wireinject
// +build wireinject

package di

import (
	"github.com/didiegovieira/go-benchmark-api/internal/infrastructure/api"

	"github.com/didiegovieira/go-benchmark-api/internal/test"
	"github.com/google/wire"
	"go.uber.org/mock/gomock"
)

var wireApiSet = wire.NewSet(
	provideMongoDbClient,
	provideApiServer,

	provideBenchmarckRepositoryMongodb,

	provideTimeCalculateUseCase,
	providePostSortingAlgorithmUseCase,

	apiMiddlewaresSet,
	apiHandlersSet,
	wire.Struct(new(api.Application), "*"),
)

var wireTestSet = wire.NewSet(
	provideMongoDbClient,
	provideApiServer,

	provideBenchmarckRepositoryMongodb,

	provideTimeCalculateUseCase,
	providePostSortingAlgorithmUseCase,

	apiMiddlewaresSet,
	apiHandlersSet,

	wire.Struct(new(api.Application), "*"),
	wire.Struct(new(test.Application), "*"),
)

func InitializeApi() (*api.Application, func(), error) {
	wire.Build(wireApiSet)
	return &api.Application{}, func() {}, nil
}

func InitializeTests(mockCtrl *gomock.Controller) (*test.Application, func(), error) {
	wire.Build(wireTestSet)

	return &test.Application{}, func() {}, nil
}
