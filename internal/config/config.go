package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	AppPort       string `env:"APP_PORT" envDefault:":8080"`
	RaribleAPIURL string `env:"RARIBLE_API_URL" envDefault:"https://api.rarible.org/v0.1"`
	RaribleAPIKey string `env:"RARIBLE_API_KEY" envDefault:""`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	return cfg, nil
}
