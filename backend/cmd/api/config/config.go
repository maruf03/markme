package api

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type AppConfig struct {
	Environment string `env:"ENVIRONMENT" envDefault:"prod"`
}

type APIConfig struct {
	AllowedOrigins []string      `env:"ALLOWED_ORIGINS,required"`
	RequestTimeout time.Duration `env:"REQUEST_TIMEOUT"          envDefault:"30"`
}

type ServerConfig struct {
	Host    string        `env:"HOST"    envDefault:"0.0.0.0"`
	Port    int32         `env:"PORT"    envDefault:"8000"`
	Timeout time.Duration `env:"TIMEOUT" envDefault:"5"`
}

type DBConfig struct {
	Host     string `env:"HOST"              envDefault:"localhost"`
	Port     int32  `env:"PORT"              envDefault:"5432"`
	Name     string `env:"NAME,required"`
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD,required"`
}

type Config struct {
	AppConfig
	APIConfig
	ServerConfig
	DBConfig DBConfig `envPrefix:"DB_"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
