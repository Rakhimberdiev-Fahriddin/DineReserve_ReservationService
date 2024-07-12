package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	DB_HOST     string
	DB_PORT     int
	DB_USERNAME string
	DB_DATABASE string
	DB_PASSWORD string
	GRPC_PORT   string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	cfg := Config{}
	cfg.DB_HOST = cast.ToString(Coalesce("DB_HOST", "localhost"))
	cfg.DB_PORT = cast.ToInt(Coalesce("DB_PORT", 5432))
	cfg.DB_USERNAME = cast.ToString(Coalesce("DB_USERNAME", "postgres"))
	cfg.DB_DATABASE = cast.ToString(Coalesce("DB_DATABASE", "reservation_service"))
	cfg.DB_PASSWORD = cast.ToString(Coalesce("DB_PASSWORD", "03212164"))
	cfg.GRPC_PORT = cast.ToString(Coalesce("GRPC_PORT", ":50051"))

	return cfg
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
