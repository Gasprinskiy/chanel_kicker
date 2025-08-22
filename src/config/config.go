package config

import (
	"os"
)

// Config структура для хранения переменных окружения
type Config struct {
	PostgresURL string
	GRPCAddr    string
}

// NewConfig загружает переменные из .env и возвращает структуру Config
func NewConfig() *Config {
	return &Config{
		PostgresURL: os.Getenv("POSTGRES_URL"),
		GRPCAddr:    os.Getenv("GRPC_ADDR"),
	}
}
