package middleware

import (
	"net/http"

	"github.com/didiegovieira/go-benchmark-api/internal/application/dto"
	"github.com/didiegovieira/go-benchmark-api/pkg/api"
	appErr "github.com/didiegovieira/go-benchmark-api/pkg/error"
	"github.com/gin-gonic/gin"
)

type RequestValidation struct {
	Presenter api.Presenter
}

func (v RequestValidation) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		var request dto.RequestInput

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !request.Validate() {
			err = dto.ErrEmptyArray
			v.Presenter.Error(c, appErr.NewHttp(http.StatusConflict, err.Error()))
		}

		c.Next()
	}

}
