package settings

import (
	"os"
	"testing"
	"time"

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

func TestInit(t *testing.T) {
	os.Setenv("MONGODB_URI", "mongodb://localhost:27017")
	os.Setenv("HTTP_SERVER_PORT", ":8080")
	os.Setenv("HTTP_SERVER_READ_TIMEOUT", "10s")
	os.Setenv("HTTP_SERVER_WRITE_TIMEOUT", "10s")

	Init()

	expectedDatabaseConnection := "mongodb://localhost:27017"
	if Settings.Database.Connection != expectedDatabaseConnection {
		t.Errorf("Expected database connection %s, got %s", expectedDatabaseConnection, Settings.Database.Connection)
	}

	expectedHttpServerPort := ":8080"
	if Settings.HttpServer.Port != expectedHttpServerPort {
		t.Errorf("Expected HTTP server port %s, got %s", expectedHttpServerPort, Settings.HttpServer.Port)
	}

	expectedReadTimeout := 10 * time.Second
	if Settings.HttpServer.ReadTimeout != expectedReadTimeout {
		t.Errorf("Expected read timeout %s, got %s", expectedReadTimeout, Settings.HttpServer.ReadTimeout)
	}

	expectedWriteTimeout := 10 * time.Second
	if Settings.HttpServer.WriteTimeout != expectedWriteTimeout {
		t.Errorf("Expected write timeout %s, got %s", expectedWriteTimeout, Settings.HttpServer.WriteTimeout)
	}
}
