package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"split-bill-backend/internal/config"
	"split-bill-backend/internal/storage"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo},
			),
		)
	default:
		// По умолчанию используем текстовый логгер
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo},
			),
		)
	}

	return log
}

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("Starting app", slog.String("env", cfg.Env))
	log.Debug("Debud messages are enabled")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	cfg.Client = storage.NewConnection(ctx, cfg)

	<-ctx.Done()
}
