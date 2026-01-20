package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env         string        `mapstructure:"ENV"`
	HTTPAddr    string        `mapstructure:"HTTP_ADDR"`
	ReadTimeout time.Duration `mapstructure:"READ_TIMEOUT"`
}

func Load() (Config, error) {
	viper.SetDefault("ENV", "dev")
	viper.SetDefault("HTTP_ADDR", ":8080")
	viper.SetDefault("READ_TIMEOUT", 10*time.Second)

	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	// If someone sets READ_TIMEOUT as "10s" in env, viper can treat it as string.
	// Keep it simple: allow Go duration strings via env if needed later.

	return cfg, nil
}
