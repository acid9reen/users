package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `env-prefix:"USERS_APP__"`
		HTTP `env-prefix:"USERS_HTTP__"`
	}
	App struct {
		Name    string `env-required:"true" env:"NAME"`
		Version string `env-required:"true" env:"VERSION"`
	}
	HTTP struct {
		Port string `env-required:"true" env:"PORT"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	envErr := cleanenv.ReadEnv(cfg)
	if envErr == nil {
		return cfg, nil
	}

	envFileErr := cleanenv.ReadConfig(".env", cfg)
	if envFileErr != nil {
		return nil, fmt.Errorf("config error: %w", envFileErr)
	}

	return cfg, nil
}
