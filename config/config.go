package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Postgres PostgresConfig
	Server   ServerConfig
}

type PostgresConfig struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_NAME     string
	DB_PASSWORD string
}

type ServerConfig struct {
	USER_PORT  string
	ITEM_PORT  string
	REDIS_PORT string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error while loading .env file: %v", err)
	}

	return &Config{
		Postgres: PostgresConfig{
			DB_HOST:     cast.ToString(coalesce("DB_HOST", "localhost")),
			DB_PORT:     cast.ToString(coalesce("DB_PORT", "5432")),
			DB_USER:     cast.ToString(coalesce("DB_USER", "postgres")),
			DB_NAME:     cast.ToString(coalesce("DB_NAME", "users")),
			DB_PASSWORD: cast.ToString(coalesce("DB_PASSWORD", "root")),
		},
		Server: ServerConfig{
			USER_PORT: cast.ToString(coalesce("USER_PORT", ":50051")),
			ITEM_PORT: cast.ToString(coalesce("ITEM_PORT", ":50052")),
			REDIS_PORT: cast.ToString(coalesce("REDIS_PORT", "6379")),

		},
	}
}

func coalesce(key string, value interface{}) interface{} {
	val, exist := os.LookupEnv(key)
	if exist {
		return val
	}
	return value
}
