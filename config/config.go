package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port     string
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
	conf := viper.New()

	conf.AutomaticEnv()

	cfg := Config{
		Port: conf.GetString("PORT"),
		Postgres: Postgres{
			Host:     conf.GetString("POSTGRES_HOST"),
			Port:     conf.GetString("POSTGRES_PORT"),
			User:     conf.GetString("POSTGRES_USER"),
			Password: conf.GetString("POSTGRES_PASSWORD"),
			Database: conf.GetString("POSTGRES_DATABASE"),
		},
	}
	return cfg
}
