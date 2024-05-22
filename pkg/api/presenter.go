package api

import "github.com/gin-gonic/gin"

type Presenter interface {
	Error(c *gin.Context, err error)
	Present(c *gin.Context, body interface{}, code int)
}
