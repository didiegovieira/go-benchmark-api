package settings

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type (
	Specification struct {
		ProjectName    string `default:"go-benchmark-api"`
		ProjectVersion string `default:"0.0.1"`
		Environment    string `envconfig:"ENVIRONMENT" default:"local"`
		Database       DatabaseSpecification
		HttpServer     HttpServerSpecification
	}

	DatabaseSpecification struct {
		Connection string `envconfig:"MONGODB_URI" required:"true"`
		DbName     string `envconfig:"MONGODB_DATABASE" default:"benchmark"`
	}

	HttpServerSpecification struct {
		Port         string        `envconfig:"HTTP_SERVER_PORT" default:":3000"`
		ReadTimeout  time.Duration `envconfig:"HTTP_SERVER_READ_TIMEOUT" default:"15s"`
		WriteTimeout time.Duration `envconfig:"HTTP_SERVER_WRITE_TIMEOUT" default:"15s"`
	}
)

var Settings Specification

func Init() {
	if err := envconfig.Process("", &Settings); err != nil {
		panic(err.Error())
	}
}

func (s *Specification) IsProduction() bool {
	return s.Environment == "production"
}

func (s *Specification) IsLocal() bool {
	return s.Environment == "local"
}
