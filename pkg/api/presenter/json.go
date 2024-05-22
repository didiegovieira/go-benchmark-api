package presenter

import (
	"net/http"

	appErr "github.com/didiegovieira/go-benchmark-api/pkg/errors"
	"github.com/gin-gonic/gin"
)

type Json struct{}

func NewJson() Json {
	return Json{}
}

func (j Json) Error(c *gin.Context, err error) {
	code := http.StatusInternalServerError
	response := gin.H{"error": err.Error()}

	switch e := err.(type) {

	case *appErr.Validation:
		code = http.StatusBadRequest
		response["messages"] = e.Errors

	case *appErr.Http:
		code = e.Code

	}

	c.AbortWithStatusJSON(code, response)
}

func (j Json) Present(c *gin.Context, body interface{}, code int) {
	c.JSON(code, body)
}
