package app

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/JPBoshoff/PsychApp/services/api/internal/config"
)

type App struct {
	cfg    config.Config
	logger *zap.Logger
	server *http.Server
}

func New(cfg config.Config, logger *zap.Logger) *App {
	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           NewRouter(),
		ReadHeaderTimeout: cfg.ReadTimeout,
	}

	return &App{
		cfg:    cfg,
		logger: logger,
		server: srv,
	}
}

func (a *App) Start() error {
	a.logger.Info("starting api", zap.String("addr", a.cfg.HTTPAddr), zap.String("env", a.cfg.Env))
	return a.server.ListenAndServe()
}

func (a *App) Stop(ctx context.Context) error {
	a.logger.Info("stopping api")
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return a.server.Shutdown(shutdownCtx)
}
