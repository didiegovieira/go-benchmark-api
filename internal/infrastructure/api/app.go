package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/didiegovieira/go-benchmark-api/internal/infrastructure/api/handler"
	"github.com/didiegovieira/go-benchmark-api/internal/infrastructure/api/middleware"
	"github.com/didiegovieira/go-benchmark-api/internal/settings"
	"github.com/didiegovieira/go-benchmark-api/pkg/api"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Server api.Server[*gin.Engine]

	HealthHandler               *handler.Health
	PostSortingAlgorithmHandler *handler.PostSortingAlgorithm

	MiddlewareCors *middleware.Cors
}

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Error loading .env file")
	// }

	settings.Init()
}

func (a *Application) Start() {
	a.SetupRoutes()

	ctx := context.Background()
	quitSig := make(chan os.Signal, 1)
	signal.Notify(quitSig, os.Interrupt)

	go func() {
		select {
		case <-quitSig:
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			if err := a.Server.Shutdown(ctx); err != nil {
				fmt.Printf("couldn't shutdown Server, err: %v", err)
			}
			return
		case <-ctx.Done():
			return
		}
	}()

	if err := a.Server.Start(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server failed to start: %v", err)
		}
		fmt.Printf("shutting down Server...")
	}
}
