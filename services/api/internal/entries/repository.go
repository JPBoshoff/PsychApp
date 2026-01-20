package entries

import "context"

// EntryRepository defines persistence behavior for journal entries.
// Implementations: memory (dev), postgres (prod).
type EntryRepository interface {
	Create(ctx context.Context, e StoredEntry) (StoredEntry, error)
	GetByID(ctx context.Context, id string) (StoredEntry, bool, error)
	
	// ListRecent returns the most recent entries, newest first.
	ListRecent(ctx context.Context, limit int) ([]StoredEntry, error)
}
