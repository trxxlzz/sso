package main

import (
	"fmt"
	"log/slog"
	"os"
	"sso/internal/app"
	_ "sso/internal/app/grpc"
	"sso/internal/config"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
	EnvDev   = "dev"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	log := SetupLogger(cfg.Env)

	log.Info("starting application")

	application := app.New(log, cfg.GRPC.Port, cfg.Storage_path, cfg.TokenTTl)

	application.GRPCSrv.MustRun()

}
func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case EnvLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case EnvDev:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
