package config

import (
	"log/slog"

	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Env struct {
	DB_NAME     string `env:"DB_NAME"`
	DB_USERNAME string `env:"DB_USERNAME"`
	DB_PASSWORD string `env:"DB_PASSWORD"`
	DB_PORT     int    `env:"DB_PORT"`
	DB_HOST     string `env:"DB_HOST"`
	IpAddress   string `env:"IP_ADDRESS"`
	API_PORT    int    `env:"API_PORT"`
}

type Config struct {
	Env    Env
	Client *pgxpool.Pool
}

var config Config

func GetConfig() *Config {
	config.Env = *getEnv()

	return &config
}

func getEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		slog.Warn("No .env file found, using system environment variables")
	}

	var cfg Env
	err = env.Parse(&cfg)
	if err != nil {
		slog.Error("Failed to parse env: %v", err)
		panic(err)
	}

	return &cfg
}

// func MustLoad() *Config {
// 	err := godotenv.Load()
// 	if err != nil {
// 		slog.Warn("no env")
// 	}

// 	configPath := os.Getenv("CONFIG_PATH")
// 	if configPath == "" {
// 		log.Fatal("Fail load cfg")
// 	}

// 	if _, err := os.Stat(configPath); os.IsNotExist(err) {
// 		log.Fatalf("config file %s does not exist", configPath)
// 	}

// 	var cfg Config

// 	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
// 		log.Fatalf("cannot read config: %s", err)
// 	}

// 	// // Подставляем переменные окружения вручную
// 	// cfg.Env.DB_HOST = os.Getenv("DB_HOST")
// 	// cfg.Env.DB_PORT = os.Getenv("DB_PORT")
// 	// cfg.Env.DB_USERNAME = os.Getenv("DB_USERNAME")
// 	// cfg.Env.DB_PASSWORD = os.Getenv("DB_PASSWORD")
// 	// cfg.Env.DB_NAME = os.Getenv("DB_NAME")
// 	// cfg.Env.API_PORT = os.Getenv("API_PORT")

// 	return &cfg
// }
