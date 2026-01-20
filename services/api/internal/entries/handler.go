package entries

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

type CreateEntryRequest struct {
	// In MVP, keep this minimal. We'll extend later.
	Text     string            `json:"text"`
	Source   string            `json:"source,omitempty"`   // e.g. "daily_checkin", "open_journal"
	Metadata map[string]string `json:"metadata,omitempty"` // optional key/value
}

type CreateEntryResponse struct {
	EntryID    string                 `json:"entry_id"`
	CreatedAt  string                 `json:"created_at"`
	Analysis   map[string]any         `json:"analysis"`
	RequestID  string                 `json:"request_id,omitempty"`
	MockNotice string                 `json:"mock_notice,omitempty"`
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
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
	entryID := "entry_" + now.Format("20060102_150405")

	resp := CreateEntryResponse{
		EntryID:   entryID,
		CreatedAt: now.Format(time.RFC3339),
		RequestID: middleware.GetReqID(r.Context()),
		MockNotice: "mock analysis - agent service not wired yet",
		Analysis: map[string]any{
			"quadrant_distribution": map[string]float64{
				"UL": 0.40,
				"UR": 0.20,
				"LL": 0.15,
				"LR": 0.25,
			},
			"themes": []string{"work_pressure", "self_criticism", "need_for_rest"},
			"mirror_reflection": map[string]any{
				"summary": "It sounds like today carried a lot of pressure and self-judgment. You’re noticing how that shows up internally, and you’re also aware your body is asking for rest. There may be a tension between what you feel you must do and what you can sustainably carry.",
				"clarifying_questions": []string{
					"What part of today felt most “non-negotiable”, and who decided that?",
					"When did you first notice the body signal that you needed rest?",
				},
			},
			"safety": map[string]any{
				"risk":               "none",
				"signals":            []string{},
				"recommended_action": "normal_flow",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}
