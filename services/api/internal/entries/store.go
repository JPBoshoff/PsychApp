package entries

import (
	"sync"
	"time"
)

type StoredEntry struct {
	EntryID   string
	CreatedAt time.Time
	Text      string
	Source    string
	Metadata  map[string]string
	Analysis  map[string]any
}

type InMemoryStore struct {
	mu   sync.RWMutex
	data map[string]StoredEntry
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		data: make(map[string]StoredEntry),
	}
}

func (s *InMemoryStore) Put(e StoredEntry) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[e.EntryID] = e
}

func (s *InMemoryStore) Get(id string) (StoredEntry, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e, ok := s.data[id]
	return e, ok
}
