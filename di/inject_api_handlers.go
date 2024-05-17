package di

import (
	"net/http"

	"github.com/didiegovieira/go-benchmark-api/internal/infrastructure/api/handler"
	"github.com/didiegovieira/go-benchmark-api/internal/settings"
	"github.com/didiegovieira/go-benchmark-api/pkg/api"
	"github.com/didiegovieira/go-benchmark-api/pkg/api/presenter"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var apiHandlersSet = wire.NewSet(
	provideApiPresenter,
	wire.Struct(new(handler.Health), "*"),
	wire.Struct(new(handler.PostSortingAlgorithm), "*"),
)

func provideApiServer() api.Server[*gin.Engine] {
	return api.NewGinServer[*gin.Engine](&http.Server{
		Addr:         settings.Settings.HttpServer.Port,
		ReadTimeout:  settings.Settings.HttpServer.ReadTimeout,
		WriteTimeout: settings.Settings.HttpServer.WriteTimeout,
	})
}

func provideApiPresenter() api.Presenter {
	return presenter.NewJson()
}
