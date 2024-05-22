package api

import (
	"github.com/didiegovieira/go-benchmark-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	prefix = "/api"
)

func (a *Application) SetupRoutes() {

	r := a.Server.GetRouter()

	docs.SwaggerInfo.Title = "Go Benchmark API"
	docs.SwaggerInfo.BasePath = "/api"

	c := r.Group(prefix)
	{
		c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		c.GET("/health", a.HealthHandler.Handle())

		sa := c.Group("/benchmark", a.MiddlewareCors.Handle())
		{
			sa.POST("/sort", a.PostSortingAlgorithmHandler.Handle())
		}
	}
}
