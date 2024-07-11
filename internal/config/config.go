package config

import "github.com/caarlos0/env/v11"

type Config struct {
	PostgresConfig
}

type PostgresConfig struct {
	postgresUser     string `env:"POSTGRES_USER,required"`
	postgresPassword string `env:"POSTGRES_PASSWORD,required"`
	postgresPort     int    `env:"POSTGRES_PORT,required"`
	postgresDB       string `env:"POSTGRES_DB,required"`
	postgresHost     string `env:"POSTGRES_HOST,required"`
}

func MustLoad() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
