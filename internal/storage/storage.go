package storage

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"split-bill-backend/internal/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// type Storage struct {
// 	db *sql.DB
// }

func NewConnection(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	env := cfg.Database

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		env.DB_USERNAME,
		env.DB_PASSWORD,
		env.DB_HOST,
		env.DB_PORT,
		env.DB_NAME,
	)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("Unable to parse config")
	}

	// Настройки пула
	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 5 * time.Minute

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		slog.Error("ошибка при подключении к БД", slog.Any("error", err))
		panic(err)
	}

	// if err = CheckAndMigrate(conn); err != nil {
	// 	slog.Error("%v", err)
	// 	panic(err)
	// }

	return conn
}
