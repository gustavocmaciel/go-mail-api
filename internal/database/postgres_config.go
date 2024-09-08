package database

import (
	"fmt"
	"os"
)

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}
}

func (config *PostgresConfig) ConnString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.DBName)
}
