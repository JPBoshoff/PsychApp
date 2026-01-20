package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/JPBoshoff/PsychApp/services/api/internal/entries"
	"github.com/JPBoshoff/PsychApp/services/api/internal/health"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	entryStore := entries.NewInMemoryStore()
	entryServer := entries.NewServer(entryStore)

	// Routes
	r.Get("/health", health.Handler)
	r.Post("/entries", entryServer.CreateHandler)
	r.Get("/entries/{entry_id}", entryServer.GetHandler)

	return r
}
