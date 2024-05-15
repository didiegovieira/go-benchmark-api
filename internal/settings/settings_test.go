package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLocal(t *testing.T) {
	assert.False(t, Settings.IsLocal())

	Settings.Environment = "local"
	assert.True(t, Settings.IsLocal())
}

func TestIsProduction(t *testing.T) {
	assert.False(t, Settings.IsProduction())

	Settings.Environment = "production"
	assert.True(t, Settings.IsProduction())
}
