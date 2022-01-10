package config

import "github.com/caarlos0/env"

type LogConfig struct {
	Level string `env:"LOG_LEVEL" envDefault:"INFO"`
}

func NewLogConfig() (LogConfig, error) {
	var conf LogConfig

	if err := env.Parse(&conf); err != nil {
		return LogConfig{}, err
	}

	return conf, nil
}
