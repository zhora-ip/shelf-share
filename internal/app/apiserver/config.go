package apiserver

import (
	"time"
)

type Config struct {
	BindAddr     string        `toml:"bind_addr"`
	LogLevel     string        `toml:"log_level"`
	ReadTimeout  time.Duration `toml:"read_time"`
	WriteTimeout time.Duration `toml:"write_time"`
	IdleTimeout  time.Duration `toml:"idle_time"`
	DatabaseURL  string        `toml:"database_url"`
	SessionKey   string        `toml:"session_key"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:     ":8080",
		LogLevel:     "debug",
		ReadTimeout:  1,
		WriteTimeout: 1,
		IdleTimeout:  120,
	}
}
