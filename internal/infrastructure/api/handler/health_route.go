package handler

import (
	"net/http"

	"github.com/didiegovieira/go-benchmark-api/pkg/api"
	"github.com/gin-gonic/gin"
)

type Health struct {
	Presenter api.Presenter
}

// PingExample godoc
// @Summary Verify api connection
// @Schemes
// @Description do ping
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {string} true
// @Router /health [get]
func (h *Health) Handle() func(c *gin.Context) {
	return func(c *gin.Context) {
		h.Presenter.Present(c, gin.H{"ok": true}, http.StatusOK)
	}
}
