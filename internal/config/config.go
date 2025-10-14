package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type HTTPServerConfig struct {
	Address     string `yaml:"address"`
	Timeout     string `yaml:"timeout" env-deafult:"4s"`
	IdleTimeout string `yaml:"idle_timeout" env-deafult:"60s"`
}

type DatabaseConfig struct {
	DB_HOST     string `yaml:"host"`
	DB_PORT     string `yaml:"port"`
	DB_USERNAME string `yaml:"username"`
	DB_PASSWORD string `yaml:"password"`
	DB_NAME     string `yaml:"database"`
	SSLMode     string `yaml:"sslmode"`
}

type Config struct {
	Env        string           `yaml:"env" env-default:"local"`
	HTTPServer HTTPServerConfig `yaml:"http_server"`
	Database   DatabaseConfig   `yaml:"database"`
	Client     *pgxpool.Pool
}

func MustLoad() *Config {
	_ = godotenv.Load()

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("Fail load cfg")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
