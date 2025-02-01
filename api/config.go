package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	User       string
	Name       string
	Password   string
	Host       string
	Port       string
	Mode       string
	Paddle_Key string
}

func NewTestConfig() *Config {
	config := new(Config)

	config.User = "postgres"
	config.Name = "postgres"
	config.Password = "root"
	config.Host = "localhost"
	config.Port = "9000"
	config.Mode = "testing"
	config.Paddle_Key = "b20c4a0db8fbdc962a6a0d5c6dbfa8dc910f0153c1a7b1b4ab"

	return config
}

func New() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := new(Config)

	config.User = os.Getenv("DB_USER")
	config.Name = os.Getenv("DB_NAME")
	config.Password = os.Getenv("DB_PASSWORD")
	config.Host = os.Getenv("DB_HOST")
	config.Port = os.Getenv("DB_PORT")
	config.Paddle_Key = os.Getenv("PADDLE_API_KEY")
	config.Mode = "DEV"

	return config
}
