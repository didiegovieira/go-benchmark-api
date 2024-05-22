package test

import (
	"net/http/httptest"

	"github.com/didiegovieira/go-benchmark-api/internal/application/repository"
	"github.com/didiegovieira/go-benchmark-api/internal/infrastructure/api"
	"github.com/didiegovieira/go-benchmark-api/internal/settings"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/mock/gomock"
)

type Application struct {
	Api *api.Application

	BenchmarkRepository repository.BenchmarkRepository
	MongoClient         *mongo.Client

	MockCtrl *gomock.Controller

	ApiUrl    string           `wire:"-"`
	ApiServer *httptest.Server `wire:"-"`
}

func (a *Application) RunApiServer() *httptest.Server {
	a.Api.SetupRoutes()

	a.ApiServer = httptest.NewServer(a.Api.Server.GetRouter())
	a.ApiUrl = a.ApiServer.URL + "/api/v2"

	return a.ApiServer
}

func (a *Application) ApiCleanup() {
	a.ApiServer.Close()
}

func (a *Application) ClearDatabase() {
	db := a.MongoClient.Database(settings.Settings.Database.DbName)
	_ = db.Collection("sequences").Drop(nil)
}
