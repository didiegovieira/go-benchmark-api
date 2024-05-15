package handler

import (
	"net/http"

	"github.com/didiegovieira/go-benchmark-api/pkg/api"
	"github.com/gin-gonic/gin"
)

type Health struct {
	Presenter api.Presenter
}

func (h *Health) Handle() func(c *gin.Context) {
	return func(c *gin.Context) {
		h.Presenter.Present(c, gin.H{"ok": true}, http.StatusOK)
	}
}
