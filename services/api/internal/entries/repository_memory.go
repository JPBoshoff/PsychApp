package entries

import (
	"context"
	"sort"
	"sync"
)

type MemoryRepository struct {
	mu   sync.RWMutex
	data map[string]StoredEntry
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data: make(map[string]StoredEntry),
	}
}

func (r *MemoryRepository) Create(ctx context.Context, e StoredEntry) (StoredEntry, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[e.EntryID] = e
	return e, nil
}

func (r *MemoryRepository) GetByID(ctx context.Context, id string) (StoredEntry, bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	e, ok := r.data[id]
	return e, ok, nil
}

func (r *MemoryRepository) ListRecent(ctx context.Context, limit int) ([]StoredEntry, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]StoredEntry, 0, len(r.data))
	for _, e := range r.data {
		out = append(out, e)
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].CreatedAt.After(out[j].CreatedAt)
	})

	if limit <= 0 {
		limit = 50
	}
	if limit > len(out) {
		limit = len(out)
	}

	return out[:limit], nil
}
