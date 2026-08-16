package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env  string
	Port string
}

func MustLoad() *Config {

	godotenv.Load()

	env := os.Getenv("APP_ENV")
	if env == "" {
		panic("APP_ENV is required")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		panic("APP_PORT is required")
	}

	return &Config{
		Env:  env,
		Port: port,
	}
}
