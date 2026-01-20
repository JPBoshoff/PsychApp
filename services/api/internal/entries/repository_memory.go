package entries

import (
	"context"
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
