package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"log"
)

type DatabaseConfig struct {
	ConnectionString string `env:"Database__ConnectionString" envDefault:"host=localhost user=oil_refining_service_user password=password dbname=oil_refining_report_service port=5432 sslmode=disable"`
}

type HttpConfig struct {
	Port string `env:"Http__Port" envDefault:"8080"`
}

type Config struct {
	Database   DatabaseConfig
	HttpConfig HttpConfig
}

func New() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("no .env file found")
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf(".env file parse error: %v\n", err)
	}

	return cfg
}
