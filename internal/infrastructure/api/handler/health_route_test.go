package handler

import (
	"net/http"
	"testing"

	"github.com/didiegovieira/go-benchmark-api/test/mocks"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestHealth(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPresenter := mocks.NewMockApiPresenter(ctrl)

	mockPresenter.EXPECT().Present(gomock.Any(), gin.H{"ok": true}, http.StatusOK).Times(1)

	route := Health{mockPresenter}
	route.Handle()(&gin.Context{})
}
