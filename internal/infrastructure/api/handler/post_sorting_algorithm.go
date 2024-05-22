package handler

import (
	"net/http"

	"github.com/didiegovieira/go-benchmark-api/internal/application/dto"
	"github.com/didiegovieira/go-benchmark-api/internal/application/usecase"
	"github.com/didiegovieira/go-benchmark-api/pkg/api"
	"github.com/didiegovieira/go-benchmark-api/pkg/errors"
	"github.com/gin-gonic/gin"
)

type PostSortingAlgorithm struct {
	PostSortingAlgorithmUseCase usecase.PostSortingAlgorithm
	Presenter                   api.Presenter
}

// Handle godoc
// @Summary Sort an array using the specified algorithm
// @Description Sorts an array based on the input sorting algorithm provided in the request body
// @Tags Sorting Algorithms
// @Accept  json
// @Produce  json
// @Param sortingInput body dto.SortingInput true "Sorting input"
// @Success 200 {object} dto.SortingOutput
// @Failure 400 {object} errors.Http
// @Failure 500 {object} errors.Http
// @Router /sort [post]
func (u *PostSortingAlgorithm) Handle() func(c *gin.Context) {
	return func(c *gin.Context) {
		var request dto.SortingInput

		if err := c.ShouldBindJSON(&request); err != nil {
			u.Presenter.Error(c, errors.NewHttp(http.StatusBadRequest, err.Error()))
			return
		}

		b, err := u.PostSortingAlgorithmUseCase.Execute(c.Request.Context(), request)
		if err != nil {
			u.Presenter.Error(c, errors.NewHttp(http.StatusInternalServerError, err.Error()))
			return
		}

		var response dto.SortingOutput

		u.Presenter.Present(c, response.NewSortingOutput(*b), http.StatusOK)
	}
}
