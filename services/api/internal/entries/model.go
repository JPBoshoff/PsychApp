package entries

import "time"

// StoredEntry represents a persisted journal entry.
// This is the canonical domain model for entries.
type StoredEntry struct {
	EntryID   string
	CreatedAt time.Time
	Text      string
	Source    string
	Metadata  map[string]string
	Analysis  map[string]any
}
