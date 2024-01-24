package config

import "github.com/ilyakaznacheev/cleanenv"

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
	return &cfg, nil
}
