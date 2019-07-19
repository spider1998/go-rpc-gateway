package app

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug    bool   `json:"debug"`
	HTTPAddr string `json:"http_addr" default:":8080"`
	Mysql    string `json:"mysql" default:"root:123456@tcp(localhost:3306)/person?charset=utf8mb4"`
	Redis    string `json:"redis" default:"localhost:6379"`
	Version  string `json:"version" default:"0.0.1"`
}

func NewConfig() (Config, error) {
	godotenv.Load()
	var config Config
	err := envconfig.Process("", &config)
	return config, err
}
