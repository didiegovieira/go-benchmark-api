package settings

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	Specification struct {
		ProjectName    string `default:"go-benchmark-api"`
		ProjectVersion string `default:"0.0.1"`
		Environment    string `envconfig:"ENVIRONMENT" default:"local"`
		Database       DatabaseSpecification
	}

	DatabaseSpecification struct {
		Connection string `envconfig:"MONGODB_URI" required:"true"`
		DbName     string `envconfig:"DB_NAME" default:"benchmark"`
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
