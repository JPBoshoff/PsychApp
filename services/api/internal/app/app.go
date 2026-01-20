package app

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/JPBoshoff/PsychApp/services/api/internal/config"
	"github.com/JPBoshoff/PsychApp/services/api/internal/entries"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	cfg    config.Config
	logger *zap.Logger
	server *http.Server
}

func New(cfg config.Config, logger *zap.Logger) *App {
	var entryRepo entries.EntryRepository

	if cfg.RepoDriver == "postgres" {
		pool, err := pgxpool.New(context.Background(), cfg.PostgresDSN)
		if err != nil {
			logger.Fatal("failed to create postgres pool", zap.Error(err))
		}
		// Optionally ping
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := pool.Ping(ctx); err != nil {
			logger.Fatal("failed to ping postgres", zap.Error(err))
		}

		entryRepo = entries.NewPostgresRepository(pool)
	} else {
		entryRepo = entries.NewMemoryRepository()
	}

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           NewRouter(entryRepo),
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
