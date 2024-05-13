package api

import (
	"fmt"
	"io"

	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/didiegovieira/go-benchmark-api/pkg/route"
	"github.com/didiegovieira/go-benchmark-api/tests/mocks"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const testServerPort = "5010"

func mockTestRouteInterface(t *testing.T, response route.Route, times int) route.RouteInterface {
	ctrl := gomock.NewController(t)
	mock := mocks.NewMockRouteInterface(ctrl)
	mock.EXPECT().GetRoute().Return(response).Times(times)
	return mock
}

func TestWebServerNewServer(t *testing.T) {

	s := NewServer(testServerPort)
	assert.Equal(t, testServerPort, s.WebServerPort)
	assert.NotNil(t, s.Gin)
}

func TestWebServerRegisterRoutes(t *testing.T) {
	mockRoute := route.Route{
		Path:     "/",
		Method:   "GET",
		Handlers: []gin.HandlerFunc{func(c *gin.Context) {}},
	}
	handlers := 1

	s := NewServer(testServerPort)
	mockedhandler := mockTestRouteInterface(t, mockRoute, handlers)
	routes := []route.RouteInterface{mockedhandler}
	s.RegisterRoutes(routes)
	registeredRoutes := s.Gin.Routes()
	assert.Equal(t, handlers, len(registeredRoutes))

}

func TestWebServerStart(t *testing.T) {
	handlers := 1
	mockRoute := route.Route{
		Path:     "/teste",
		Method:   "GET",
		Handlers: []gin.HandlerFunc{func(c *gin.Context) { c.Data(http.StatusOK, "text/plain", []byte("up")) }},
	}

	mockedhandler := mockTestRouteInterface(t, mockRoute, handlers)
	routes := []route.RouteInterface{mockedhandler}

	s := NewServer(testServerPort)
	s.RegisterRoutes(routes)

	go s.Start()
	//defer s.server.Shutdown(context.Background())

	timeout := 10 * time.Millisecond
	iterations := 0
	up := false
	for !up && iterations < 200 {
		iterations++
		_, erro := net.DialTimeout("tcp", "localhost:"+testServerPort, timeout)
		up = (erro == nil)
	}

	requestURL := fmt.Sprintf("http://localhost:%s/teste", testServerPort)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	responseData, _ := io.ReadAll(res.Body)

	assert.Equal(t, []byte("up"), responseData)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestWebServerStartError(t *testing.T) {
	defer func() { log.StandardLogger().ExitFunc = nil }()
	var fatal bool
	log.StandardLogger().ExitFunc = func(int) { fatal = true }

	s := NewServer("UMPALUMPA")
	go s.Start()
	iterations := 0
	for !fatal && iterations < 200 {
		iterations++
		time.Sleep(10 * time.Millisecond)
	}
	assert.True(t, fatal)
}
