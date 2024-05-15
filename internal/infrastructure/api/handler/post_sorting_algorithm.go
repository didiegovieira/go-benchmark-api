package handler

import (
	"net/http"

	"github.com/didiegovieira/go-benchmark-api/internal/application/dto"
	"github.com/didiegovieira/go-benchmark-api/internal/application/usecase"

	"github.com/gin-gonic/gin"
)

type PostSortingAlgorithm struct {
	PostSortingAlgorithmUseCase usecase.PostSortingAlgorithmUseCase
}

func (u *PostSortingAlgorithm) Handle() func(c *gin.Context) {
	return func(c *gin.Context) {
		var request dto.RequestInput

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		b, err := u.PostSortingAlgorithmUseCase.Execute(c.Request.Context(), request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Sorting Algorithm",
			"status":  200,
			"result": gin.H{
				"Fastest: ": b.Fast.Duration.String(),
				"Slowest: ": b.Slow.Duration.String(),
			},
			"data": b,
		})
	}
}
