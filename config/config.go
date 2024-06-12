package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Postgres Postgres
}

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load(path string) Config {
	err := godotenv.Load(path + "/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	viper.AutomaticEnv()

	cfg := Config{
		Postgres: Postgres{
			Host:     viper.GetString("POSTGRES_HOST"),
			Port:     viper.GetString("POSTGRES_PORT"),
			User:     viper.GetString("POSTGRES_USER"),
			Password: viper.GetString("POSTGRES_PASSWORD"),
			Database: viper.GetString("POSTGRES_DATABASE"),
		},
	}
	return cfg
}
