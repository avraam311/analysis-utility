package main

import (
	"context"
	"errors"
	"os/signal"
	"syscall"
	"time"

	handlerAn "github.com/avraam311/analysis-utility/internal/api/http/handlers/analysis"
	"github.com/avraam311/analysis-utility/internal/api/http/server"
	"github.com/avraam311/analysis-utility/internal/infra/config"
	"github.com/avraam311/analysis-utility/internal/infra/logger"
	serviceAn "github.com/avraam311/analysis-utility/internal/service/analysis"
)

const (
	configFilePath = "config/local.yaml"
	envFilePath    = ".env"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger.Init()
	cfg := config.New()
	if err := cfg.LoadEnvFiles(envFilePath); err != nil {
		logger.Logger.Fatal().Err(err).Msg("failed to load env file")
	}
	cfg.EnableEnv("")
	if err := cfg.LoadConfigFiles(configFilePath); err != nil {
		logger.Logger.Fatal().Err(err).Msg("failed to load config file")
	}

	srvsAn := serviceAn.New()
	handAn := handlerAn.New(srvsAn)

	router := server.NewRouter(cfg, handAn)
	srv := server.NewServer(cfg.GetString("server.port"), router)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Logger.Fatal().Err(err).Msg("failed to run server")
		}
	}()
	logger.Logger.Info().Msg("server is running")

	<-ctx.Done()
	logger.Logger.Info().Msg("shutdown signal received")

	shutdownCtx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	logger.Logger.Info().Msg("shutting down")
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Logger.Error().Err(err).Msg("failed to shutdown server")
	}
	if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
		logger.Logger.Info().Msg("timeout exceeded, forcing shutdown")
	}
}
