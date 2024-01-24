package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	DbConfig   DbConfig   `json:"db_config"`
	HttpServer HttpServer `json:"http_server"`
}

type DbConfig struct {
	DbHost     string `json:"db_host" env:"DB_HOST"`
	DbPort     string `json:"db_port" env:"DB_PORT"`
	DbUser     string `json:"db_user" env:"DB_USER"`
	DbPassword string `json:"db_password" env:"DB_PASSWORD"`
	DbName     string `json:"db_name" env:"DB_NAME"`
	DbSllMode  string `json:"db_sll_mode" env:"DB_SLL_MODE"`
}

type HttpServer struct {
	ServerHost string `json:"server_host" env:"SERVER_HOST"`
	ServerPort string `json:"server_port" env:"SERVER_PORT"`
	AppName    string `json:"app_name" env:"APP_NAME"`
	AppHeader  string `json:"app_header" env:"APP_HEADER"`
}

func GetConfig() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig("../../.env", &cfg)
	if err != nil {
		return nil, err
	}
	cfg = Config{
		DbConfig: DbConfig{
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbName:     os.Getenv("DB_NAME"),
			DbSllMode:  os.Getenv("DB_SLL_MODE"),
		},
		HttpServer: HttpServer{
			ServerHost: os.Getenv("SERVER_HOST"),
			ServerPort: os.Getenv("SERVER_PORT"),
			AppName:    os.Getenv("APP_NAME"),
			AppHeader:  os.Getenv("APP_HEADER"),
		},
	}
	return &cfg, nil
}
