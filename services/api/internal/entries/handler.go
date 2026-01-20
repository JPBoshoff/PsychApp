package entries

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type CreateEntryRequest struct {
	Text     string            `json:"text"`
	Source   string            `json:"source,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

type CreateEntryResponse struct {
	EntryID    string         `json:"entry_id"`
	CreatedAt  string         `json:"created_at"`
	Text       string         `json:"text,omitempty"`
	Analysis   map[string]any `json:"analysis"`
	RequestID  string         `json:"request_id,omitempty"`
	MockNotice string         `json:"mock_notice,omitempty"`
}

type GetEntryResponse struct {
	EntryID   string         `json:"entry_id"`
	CreatedAt string         `json:"created_at"`
	Text      string         `json:"text"`
	Source    string         `json:"source,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	Analysis  map[string]any `json:"analysis"`
}

type Server struct {
	repo EntryRepository
	analyzer Analyzer
}

func NewServer(repo EntryRepository, analyzer Analyzer) *Server {
	return &Server{repo: repo, analyzer: analyzer}
}

func (s *Server) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateEntryRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	if req.Text == "" {
		http.Error(w, "`text` is required", http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()
	entryID := "entry_" + now.Format("20060102_150405.000")

	analysis, err := s.analyzer.Analyze(r.Context(), req.Text, req.Source, req.Metadata)
	if err != nil {
		http.Error(w, "agent error", http.StatusBadGateway)
		return
	}

	_, _ = s.repo.Create(r.Context(), StoredEntry{
	EntryID:   entryID,
	CreatedAt: now,
	Text:      req.Text,
	Source:    req.Source,
	Metadata:  req.Metadata,
	Analysis:  analysis,
})


	resp := CreateEntryResponse{
		EntryID:    entryID,
		CreatedAt:  now.Format(time.RFC3339),
		Analysis:   analysis,
		RequestID:  middleware.GetReqID(r.Context()),
		MockNotice: "mock analysis - stored in-memory (dev only)",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "entry_id")
	if id == "" {
		http.Error(w, "missing entry_id", http.StatusBadRequest)
		return
	}

	e, ok, err := s.repo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	resp := GetEntryResponse{
		EntryID:   e.EntryID,
		CreatedAt: e.CreatedAt.UTC().Format(time.RFC3339),
		Text:      e.Text,
		Source:    e.Source,
		Metadata:  e.Metadata,
		Analysis:  e.Analysis,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
