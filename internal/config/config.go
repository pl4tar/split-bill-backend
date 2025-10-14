package config

import (
	"log"
	"log/slog"
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
	Env              string `yaml:"env" env-default:"local"`
	HTTPServerConfig `yaml:"http_server"`
	DatabaseConfig   `yaml:"database"`
	Client           *pgxpool.Pool
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		slog.Warn("no env")
	}

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

	// Подставляем переменные окружения вручную
	cfg.DatabaseConfig.DB_HOST = os.Getenv("DB_HOST")
	cfg.DatabaseConfig.DB_PORT = os.Getenv("DB_PORT")
	cfg.DatabaseConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	cfg.DatabaseConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	cfg.DatabaseConfig.DB_NAME = os.Getenv("DB_NAME")

	return &cfg
}
