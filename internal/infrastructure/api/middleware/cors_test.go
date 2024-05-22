package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/didiegovieira/go-benchmark-api/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCorsMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mockPresenter := mocks.NewMockApiPresenter(ctrl)

	t.Run("should set CORS headers for all requests", func(t *testing.T) {
		router := gin.New()
		cors := Cors{Presenter: mockPresenter}
		router.Use(cors.Handle())

		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "test")
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
		assert.Equal(t, "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With", w.Header().Get("Access-Control-Allow-Headers"))
		assert.Equal(t, "POST, OPTIONS, GET, PUT, DELETE, PATCH", w.Header().Get("Access-Control-Allow-Methods"))
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "test", w.Body.String())
	})

	t.Run("should handle OPTIONS request and abort with status 204", func(t *testing.T) {
		router := gin.New()
		cors := Cors{Presenter: mockPresenter}
		router.Use(cors.Handle())

		router.POST("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "test")
		})

		req, _ := http.NewRequest(http.MethodOptions, "/test", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
		assert.Equal(t, "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With", w.Header().Get("Access-Control-Allow-Headers"))
		assert.Equal(t, "POST, OPTIONS, GET, PUT, DELETE, PATCH", w.Header().Get("Access-Control-Allow-Methods"))
		assert.Equal(t, http.StatusNoContent, w.Code)
		assert.Equal(t, "", w.Body.String())
	})
}
