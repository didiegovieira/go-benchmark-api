package middleware

import (
	"net/http"

	"github.com/didiegovieira/go-benchmark-api/internal/application/dto"
	"github.com/didiegovieira/go-benchmark-api/pkg/api"
	cer "github.com/didiegovieira/go-benchmark-api/pkg/error"
	"github.com/gin-gonic/gin"
)

type RequestValidation struct {
	Presenter api.Presenter
}

func (v RequestValidation) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		var request dto.SortingInput

		if err := c.ShouldBindJSON(&request); err != nil {
			v.Presenter.Error(c, cer.NewHttp(http.StatusBadRequest, err.Error()))
			return
		}

		if !request.Validate() {
			err = dto.ErrEmptyArray
			v.Presenter.Error(c, cer.NewHttp(http.StatusConflict, err.Error()))
			return
		}

		c.Next()
	}

}
