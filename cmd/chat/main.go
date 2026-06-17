package main

import (
	"chat/internal/app"
	"chat/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const envLocal = "local"
const envDev = "dev"
const envProd = "prod"

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	application := app.New(log, cfg.GRPCServer, cfg.Postgres)
	go func() {
		application.GRPCServer.MustRun()
	}()
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGTERM, syscall.SIGINT)
	notify := <-sign
	application.GRPCServer.Stop()
	log.Info("server stopped", slog.String("signal", notify.String()))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}))
	default:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}))
	}
	return log

}
