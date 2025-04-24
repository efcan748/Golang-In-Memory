package store

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	mu      sync.RWMutex
	Strings map[string]stringEntry
	Lists   map[string]listEntry
}

type stringEntry struct {
	Value      string
	expiration time.Time
}

type listEntry struct {
	Values     []string
	expiration time.Time
}

func New() *Store {
	return &Store{
		Strings: make(map[string]stringEntry),
		Lists:   make(map[string]listEntry),
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

	for key, entry := range s.Strings {
		if !entry.expiration.IsZero() && now.After(entry.expiration) {
			fmt.Println("Deleted String", key)
			delete(s.Strings, key)
		}
	}

	for key, entry := range s.Lists {
		if !entry.expiration.IsZero() && now.After(entry.expiration) {
			fmt.Println("Deleted List", key)
			delete(s.Lists, key)
		}
	}
}
