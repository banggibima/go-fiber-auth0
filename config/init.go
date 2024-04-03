package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	App
	Auth0
}

type App struct {
	Name string `env:"APP_NAME"`
	Port int    `env:"APP_PORT"`
}

type Auth0 struct {
	Domain   string `env:"AUTH0_DOMAIN"`
	Audience string `env:"AUTH0_AUDIENCE"`
}

func ConfigInit() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
