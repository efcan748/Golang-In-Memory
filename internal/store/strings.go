package store

import (
	"fmt"
	"time"
)

func (s *Store) Set(key, value string, ttl time.Duration) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.strings[key]
	if exists {
		return fmt.Sprintf("Key %s is already exist", key), false
	}

	expiration := time.Time{}
	if ttl > 0 {
		expiration = time.Now().Add(ttl)
	}

	s.strings[key] = stringEntry{
		value:      value,
		expiration: expiration,
	}

	return "", true
}

func (s *Store) Update(key, value string, ttl time.Duration) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry, exists := s.strings[key]
	if !exists {
		return fmt.Sprintf("key %s not found", key), false
	}

	if !entry.expiration.IsZero() && time.Now().After(entry.expiration) {
		return "entity is expired", false
	}

	expiration := time.Time{}
	if ttl > 0 {
		expiration = time.Now().Add(ttl)
	}

	entry.value = value
	entry.expiration = expiration

	s.strings[key] = entry

	return "", true
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	entry, exists := s.strings[key]
	if !exists {
		return "", false
	}

	if !entry.expiration.IsZero() && time.Now().After(entry.expiration) {
		return "", false
	}

	return entry.value, true
}

func (s *Store) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.strings[key]; exists {
		delete(s.strings, key)
		return true
	} else {
		return false
	}

}
