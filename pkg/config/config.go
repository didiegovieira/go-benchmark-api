package config

import (
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	log "github.com/sirupsen/logrus"
)

type Conf struct {
	k *koanf.Koanf
}

func (c *Conf) GetConfig(key string, defaultValue string) string {
	value := c.k.String(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func (c *Conf) GetOrPanic(key string, message string) string {
	value := c.GetConfig(key, "")
	if value == "" {
		panic(message)
	}

	return value
}

func LoadConfig() *Conf {
	var k = koanf.New(".")
	cfg := Conf{
		k: k,
	}

	if err := k.Load(file.Provider(".env"), dotenv.Parser()); err != nil {
		log.Infoln(".env file not found")
	}

	_ = k.Load(env.Provider("", ".", func(s string) string {
		return s
	}), nil)

	return &cfg
}
