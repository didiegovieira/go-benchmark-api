package strings

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestStringToInt(t *testing.T) {
	assert.Equal(t, ToInt("10"), 10)
	assert.Equal(t, ToInt("invalid"), 0)
	assert.Equal(t, ToInt("10 "), 0)
}

func TestFloatToString(t *testing.T) {
	assert.Equal(t, FloatToString(10.0), "10")
	assert.Equal(t, FloatToString(10.1), "10.1")
	assert.Equal(t, FloatToString(10.12), "10.12")
	assert.Equal(t, FloatToString(10.123), "10.123")
	assert.Equal(t, FloatToString(10.1234), "10.1234")
}
