package api

import "time"

type AppConfig struct {
	Environment string
}

type APIConfig struct {
	AllowedOrigins []string
	RequestTimeout time.Duration
}

type ServerConfig struct {
	Host    string
	Port    int32
	Timeout time.Duration
}

type DBConfig struct {
}

type LogConfig struct {
	Level  string
	Format string
}

type Config interface {
	AppConfig | APIConfig | ServerConfig | DBConfig
}

type Option[T Config] func(*T)

func New[T Config](t *T, options ...Option[T]) *T {
	for _, option := range options {
		option(t)
	}
	return t
}
