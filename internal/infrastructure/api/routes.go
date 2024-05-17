package api

const (
	baseRoutePrefix = "/api"
)

func (a *Application) SetupRoutes() {

	router := a.Server.GetRouter()

	base := router.Group(baseRoutePrefix)
	{
		base.GET("/health", a.HealthHandler.Handle())

		sortingAlgorithms := base.Group("/benchmark", a.MiddlewareValidationRequest.Handle(), a.MiddlewareCors.Handle())
		{
			sortingAlgorithms.POST("/sort", a.PostSortingAlgorithmHandler.Handle())
		}
	}
}
