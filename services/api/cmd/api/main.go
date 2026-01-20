package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/JPBoshoff/PsychApp/services/api/internal/app"
	"github.com/JPBoshoff/PsychApp/services/api/internal/config"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func() { _ = logger.Sync() }()

	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	a := app.New(cfg, logger)

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		_ = a.Stop(context.Background())
	}()

	if err := a.Start(); err != nil {
		logger.Fatal("api stopped with error", zap.Error(err))
	}
}
