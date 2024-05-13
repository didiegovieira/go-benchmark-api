package route

import "github.com/gin-gonic/gin"

type Route struct {
	Path     string
	Method   string
	Handlers []gin.HandlerFunc
}

type RouteInterface interface {
	GetRoute() Route
}
