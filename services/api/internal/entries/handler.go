package entries

import (
	"encoding/json"
	"net/http"
	"strconv"
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

type EntryListItem struct {
	EntryID   string   `json:"entry_id"`
	CreatedAt string   `json:"created_at"`
	Source    string   `json:"source,omitempty"`
	Excerpt   string   `json:"excerpt"`
	Themes    []string `json:"themes,omitempty"`
}

type ListEntriesResponse struct {
	Items []EntryListItem `json:"items"`
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
	createdAt := now.Format(time.RFC3339)

	reqID := middleware.GetReqID(r.Context())

	analysis, err := s.analyzer.Analyze(
		r.Context(),
		reqID,
		entryID,
		createdAt,
		req.Text,
		req.Source,
		req.Metadata,
	)
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

func makeExcerpt(s string, max int) string {
	if max <= 0 {
		max = 140
	}
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	return string(r[:max]) + "..."
}

func (s *Server) ListHandler(w http.ResponseWriter, r *http.Request) {
	limit := 50
	if v := r.URL.Query().Get("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			limit = n
		}
	}
	if limit <= 0 {
		limit = 50
	}
	if limit > 200 {
		limit = 200
	}

	entries, err := s.repo.ListRecent(r.Context(), limit)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	items := make([]EntryListItem, 0, len(entries))
	for _, e := range entries {
		// Try pull themes from analysis if available
		themes := []string{}
		if raw, ok := e.Analysis["themes"]; ok {
			// themes is expected to be []any when unmarshaled
			if arr, ok := raw.([]any); ok {
				for _, t := range arr {
					if s, ok := t.(string); ok {
						themes = append(themes, s)
					}
				}
			}
		}

		items = append(items, EntryListItem{
			EntryID:   e.EntryID,
			CreatedAt: e.CreatedAt.UTC().Format(time.RFC3339),
			Source:    e.Source,
			Excerpt:   makeExcerpt(e.Text, 140),
			Themes:    themes,
		})
	}

	resp := ListEntriesResponse{Items: items}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
