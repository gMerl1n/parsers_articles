package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ConfigServer
	Token    ConfigToken
	Postgres ConfigPostgres
	Redis    ConfigRedis
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

type ConfigServer struct {
	Port string
}

type ConfigToken struct {
	JWTsecret       string
	AccessTokenTTL  int
	RefreshTokenTTL int
}

func fetchConfig() error {

	configPath := filepath.Join("configs")
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}

func NewConfig() (*Config, error) {

	if err := fetchConfig(); err != nil {
		fmt.Printf("error initialization config %s", err.Error())
		return nil, err
	}

	dbr, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		return nil, err
	}

	return &Config{

		Server: ConfigServer{
			Port: viper.GetString("server.port"),
		},
		Token: ConfigToken{
			JWTsecret:       viper.GetString("token.jwt_secret"),
			AccessTokenTTL:  viper.GetInt("token.access_token_TTL"),
			RefreshTokenTTL: viper.GetInt("token.refresh_token_TTL"),
		},

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
		Loglevel: os.Getenv("LOGLEVEL"),
	}, nil
}
