package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env         string        `mapstructure:"ENV"`
	HTTPAddr    string        `mapstructure:"HTTP_ADDR"`
	ReadTimeout time.Duration `mapstructure:"READ_TIMEOUT"`
	RepoDriver  string        `mapstructure:"REPO_DRIVER"`
	PostgresDSN string `mapstructure:"POSTGRES_DSN"`
	AgentDriver string `mapstructure:"AGENT_DRIVER"` // "mock" or "python"
	AgentURL    string `mapstructure:"AGENT_URL"`    // e.g. http://127.0.0.1:8091
}

func Load() (Config, error) {
	viper.SetDefault("ENV", "dev")
	viper.SetDefault("HTTP_ADDR", ":8080")
	viper.SetDefault("READ_TIMEOUT", 10*time.Second)
	viper.SetDefault("REPO_DRIVER", "memory")
	viper.SetDefault("POSTGRES_DSN", "postgres://postgres:postgres@localhost:5432/psychapp?sslmode=disable")

	viper.SetDefault("AGENT_DRIVER", "mock")
	viper.SetDefault("AGENT_URL", "http://127.0.0.1:8091")

	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	// If someone sets READ_TIMEOUT as "10s" in env, viper can treat it as string.
	// Keep it simple: allow Go duration strings via env if needed later.

	return cfg, nil
}
