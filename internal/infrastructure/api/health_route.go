package api

import (
	"net/http"

	"github.com/didiegovieira/go-benchmark-api/pkg/route"
	"github.com/gin-gonic/gin"
)

type HealthRoute struct {
	path   string
	method string
}

func NewHealthRoute() *HealthRoute {
	return &HealthRoute{
		path:   "/health",
		method: "GET",
	}
}

func (h *HealthRoute) getHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain", []byte("OK"))
	}
}

func (h *HealthRoute) GetRoute() route.Route {
	return route.Route{
		Path:     h.path,
		Method:   h.method,
		Handlers: []gin.HandlerFunc{h.getHandler()},
	}
}
