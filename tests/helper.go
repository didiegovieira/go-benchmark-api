package tests

import (
	"testing"

	"go.uber.org/mock/gomock"
)

func Setup(t *testing.T, setEnv *map[string]string) *gomock.Controller {
	t.Setenv("ENVIRONMENT", "testing")

	if setEnv != nil {
		for key, value := range *setEnv {
			t.Setenv(key, value)
		}
	}

	return gomock.NewController(t)
}
