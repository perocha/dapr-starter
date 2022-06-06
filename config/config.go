package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config
	Config struct {
		App `yaml:"app"`
	}

	// App
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		Port    string `env-required:"true" yaml:"port"    env:"APP_PORT"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		log.Fatalf("config error: %s", err)
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
