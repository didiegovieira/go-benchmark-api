//go:build wireinject
// +build wireinject

package main

import (
	usecase "github.com/didiegovieira/go-benchmark-api/internal/application/use_case"
	"github.com/didiegovieira/go-benchmark-api/pkg/config"
	"github.com/didiegovieira/go-benchmark-api/pkg/route"
	"github.com/google/wire"
)

type DependencyContainer struct {
	Configs *config.Conf

	Routes []route.RouteInterface

	PostSortingAlgorithmUseCase usecase.PostSortingAlgorithmUseCaseInterface
	TimeCalculateUseCase        usecase.TimeCalculateUseCaseInterface
}

func newDependencyContainer(
	configs *config.Conf,

	routes []route.RouteInterface,

	postSortingAlgorithmUseCase usecase.PostSortingAlgorithmUseCaseInterface,
	timeCalculateUseCase usecase.TimeCalculateUseCaseInterface,

) DependencyContainer {

	dc := DependencyContainer{
		Configs: configs,

		Routes: routes,
	}

	return dc
}

func InitializeDependencyContainer() (DependencyContainer, error) {
	wire.Build(
		config.LoadConfig,
		clientsSet,
		repositoriesSet,
		useCasesSet,
		routesSet,
		newDependencyContainer,
	)

	return DependencyContainer{}, nil
}
