package configs

import (
	"os"
	"strconv"
)

type Config struct {
	Postgres ConfigPostgres
	Redis    ConfigRedis
	Bindaddr string
	Loglevel string
}

type ConfigPostgres struct {
	Host     string
	Port     string
	User     string
	Password string
	NameDB   string
	SSLMode  string
}

type ConfigRedis struct {
	AddrRedis     string
	PasswordRedis string
	DBRedis       int
}

func NewConfig() (*Config, error) {

	dbr, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		return nil, err
	}

	return &Config{
		Postgres: ConfigPostgres{
			User:     os.Getenv("POSTGRES_USER"),
			NameDB:   os.Getenv("POSTGRES_DB"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			SSLMode:  os.Getenv("SSLMODE"),
		},
		Redis: ConfigRedis{
			AddrRedis:     os.Getenv("REDIS_PORT"),
			PasswordRedis: os.Getenv("REDIS_PASSWORD"),
			DBRedis:       dbr,
		},
		Bindaddr: os.Getenv("BINDADDR"),
		Loglevel: os.Getenv("LOGLEVEL"),
	}, nil
}
