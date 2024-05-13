package main

import (
	usecase "github.com/didiegovieira/go-benchmark-api/internal/application/use_case"
	"github.com/didiegovieira/go-benchmark-api/internal/infrastructure/api"
	"github.com/didiegovieira/go-benchmark-api/pkg/route"
	"github.com/google/wire"
)

var routesSet = wire.NewSet(
	provideHealthRoute,
	providePostSortingAlgorithm,
	provideRoutes,
)

func provideHealthRoute() *api.HealthRoute {
	return api.NewHealthRoute()
}

func providePostSortingAlgorithm(PostSortingAlgorithmUseCase usecase.PostSortingAlgorithmUseCaseInterface) *api.PostSortingAlgorithmRoute {
	return api.NewPostSortingAlgorithmRoute(PostSortingAlgorithmUseCase)
}

func provideRoutes(
	h *api.HealthRoute,
	s *api.PostSortingAlgorithmRoute,
) []route.RouteInterface {
	return []route.RouteInterface{h, s}
}
