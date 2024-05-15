package api

const (
	baseRoutePrefix = "/api"
)

func (a *Application) SetupRoutes() {

	router := a.Server.GetRouter()

	base := router.Group(baseRoutePrefix)
	{
		base.GET("/health", a.HealthHandler.Handle())

		sortingAlgorithms := router.Group("/benchmark", a.MiddlewareValidationRequest.Handle())
		{
			sortingAlgorithms.POST("/sort")
		}
	}
}
