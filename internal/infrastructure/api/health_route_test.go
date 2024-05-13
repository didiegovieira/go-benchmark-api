package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	testHelthPath   string = "/health"
	testHelthMethod string = "GET"
)

func TestWebHealthNewHealthRoute(t *testing.T) {
	healthRoute := NewHealthRoute()

	assert.Equal(t, testHelthPath, healthRoute.path)
	assert.Equal(t, testHelthMethod, healthRoute.method)
}

func TestWebHealthgetHandler(t *testing.T) {
	mockResponse := "OK"
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	route := NewHealthRoute()
	router.GET(route.path, route.getHandler())
	req, _ := http.NewRequest(http.MethodGet, route.path, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	reponseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse, string(reponseData))
}

func TestWebHealthGetRoute(t *testing.T) {
	healthRoute := NewHealthRoute()
	route := healthRoute.GetRoute()

	assert.Equal(t, testHelthPath, route.Path)
	assert.Equal(t, testHelthMethod, route.Method)
	assert.Len(t, route.Handlers, 1)
}
