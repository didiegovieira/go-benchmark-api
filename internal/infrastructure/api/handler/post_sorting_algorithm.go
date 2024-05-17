package handler

import (
	"net/http"

	"github.com/didiegovieira/go-benchmark-api/internal/application/dto"
	"github.com/didiegovieira/go-benchmark-api/internal/application/usecase"
	"github.com/didiegovieira/go-benchmark-api/pkg/api"
	cer "github.com/didiegovieira/go-benchmark-api/pkg/error"
	"github.com/gin-gonic/gin"
)

type PostSortingAlgorithm struct {
	PostSortingAlgorithmUseCase usecase.PostSortingAlgorithm
	Presenter                   api.Presenter
}

func (u *PostSortingAlgorithm) Handle() func(c *gin.Context) {
	return func(c *gin.Context) {
		var request dto.SortingInput

		if err := c.ShouldBindJSON(&request); err != nil {
			u.Presenter.Error(c, cer.NewHttp(http.StatusBadRequest, err.Error()))
			return
		}

		b, err := u.PostSortingAlgorithmUseCase.Execute(c.Request.Context(), request)
		if err != nil {
			u.Presenter.Error(c, cer.NewHttp(http.StatusInternalServerError, err.Error()))
			return
		}

		var response dto.SortingOutput

		u.Presenter.Present(c, response.NewSortingOutput(*b), http.StatusOK)
	}
}
