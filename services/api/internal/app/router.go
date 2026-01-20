package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/JPBoshoff/PsychApp/services/api/internal/health"
	"github.com/JPBoshoff/PsychApp/services/api/internal/entries"

)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// Routes
	r.Get("/health", health.Handler)
	r.Post("/entries", entries.CreateHandler)

	return r
}
