package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/JPBoshoff/PsychApp/services/api/internal/entries"
	"github.com/JPBoshoff/PsychApp/services/api/internal/health"
)

func NewRouter(entryRepo entries.EntryRepository, analyzer entries.Analyzer) http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Request-Id"},
		ExposedHeaders:   []string{"X-Request-Id"},
		AllowCredentials: false,
		MaxAge:           300, // 5 minutes
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	entryServer := entries.NewServer(entryRepo, analyzer)

	r.Get("/health", health.Handler)
	r.Post("/entries", entryServer.CreateHandler)
	r.Get("/entries", entryServer.ListHandler)
	r.Get("/entries/{entry_id}", entryServer.GetHandler)

	return r
}
