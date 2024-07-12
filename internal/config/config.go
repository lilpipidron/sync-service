package config

import "github.com/caarlos0/env/v11"

type Config struct {
	PostgresConfig
	ServiceConfig
	KubernetesConfig
}

type KubernetesConfig struct {
	Namespace string `env:"KUBERNETES_NAMESPACE"`
}

type ServiceConfig struct {
	ServiceHost string `env:"SERVICE_HOST,required"`
	ServicePort int    `env:"SERVICE_PORT,required"`
}

type PostgresConfig struct {
	PostgresUser     string `env:"POSTGRES_USER,required"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,required"`
	PostgresPort     int    `env:"POSTGRES_PORT,required"`
	PostgresDB       string `env:"POSTGRES_DB,required"`
	PostgresHost     string `env:"POSTGRES_HOST,required"`
}

func MustLoad() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
