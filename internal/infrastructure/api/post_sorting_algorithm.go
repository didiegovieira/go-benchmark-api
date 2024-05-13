package api

import (
	"net/http"

	"github.com/didiegovieira/go-benchmark-api/internal/application/dto"
	usecase "github.com/didiegovieira/go-benchmark-api/internal/application/use_case"
	"github.com/didiegovieira/go-benchmark-api/pkg/route"

	"github.com/gin-gonic/gin"
)

type PostSortingAlgorithmRoute struct {
	path                        string
	method                      string
	PostSortingAlgorithmUseCase usecase.PostSortingAlgorithmUseCaseInterface
}

func NewPostSortingAlgorithmRoute(postSortingAlgorithmUseCase usecase.PostSortingAlgorithmUseCaseInterface) *PostSortingAlgorithmRoute {
	return &PostSortingAlgorithmRoute{
		path:                        "/sorting-algorithm",
		method:                      "POST",
		PostSortingAlgorithmUseCase: postSortingAlgorithmUseCase,
	}
}

func (u *PostSortingAlgorithmRoute) getHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var request dto.Request

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u.PostSortingAlgorithmUseCase.Execute(request.Arr)

		c.JSON(http.StatusOK, gin.H{
			"message": "Sorting Algorithm",
			"status":  200,
		})
	}
}

func (u *PostSortingAlgorithmRoute) GetRoute() route.Route {
	return route.Route{
		Path:     u.path,
		Method:   u.method,
		Handlers: []gin.HandlerFunc{u.getHandler()},
	}
}
