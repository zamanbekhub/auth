package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"sync"
)

type (
	Config struct {
		Service  *Service
		Database *Database
	}

	Service struct {
		Port string `envconfig:"port" default:"8000"`
	}

	Database struct {
		PostgreDSN string `envconfig:"POSTGRE_DSN" required:"true"`
	}
)

var (
	once   sync.Once
	config *Config
)

// GetConfig Загружает конфиг из .env файла и возвращает объект конфигурации
// В случае, если не передать параметр `envfiles`, берется `.env` файл из корня проекта
func GetConfig(envfiles ...string) (*Config, error) {
	var err error
	once.Do(func() {
		_ = godotenv.Load(envfiles...)

		var c Config
		err = envconfig.Process("", &c)
		if err != nil {
			err = fmt.Errorf("error parse config from env variables: %w\n", err)
			return
		}

		config = &c
	})

	return config, err
}
