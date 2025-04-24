package store

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	mu      sync.RWMutex
	strings map[string]stringEntry
	lists   map[string]listEntry
}

type stringEntry struct {
	value      string
	expiration time.Time
}

type listEntry struct {
	values     []string
	expiration time.Time
}

func New() *Store {
	return &Store{
		strings: make(map[string]stringEntry),
		lists:   make(map[string]listEntry),
	}
}

func (s *Store) StartCleanup(interval time.Duration) *Store {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			s.cleanupExpired()
		}
	}()
	return s
}

func (s *Store) cleanupExpired() {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()

	for key, entry := range s.strings {
		if !entry.expiration.IsZero() && now.After(entry.expiration) {
			fmt.Println("Deleted String", key)
			delete(s.strings, key)
		}
	}

	for key, entry := range s.lists {
		if !entry.expiration.IsZero() && now.After(entry.expiration) {
			fmt.Println("Deleted List", key)
			delete(s.lists, key)
		}
	}
}
