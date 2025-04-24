package store

import (
	"fmt"
	"time"
)

func (s *Store) LPush(key string, ttl time.Duration, values ...string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	expiration := time.Time{}
	if ttl > 0 {
		expiration = time.Now().Add(ttl)
	}

	entry := s.Lists[key]
	fmt.Println(entry.Values, entry.expiration, key)

	s.Lists[key] = listEntry{
		Values:     append(values, entry.Values...),
		expiration: expiration,
	}
	fmt.Println(s.Lists[key].Values, s.Lists[key].expiration, key, "Updated")

	return "", true

}

func (s *Store) Pop(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	entry, exists := s.Lists[key]

	if !exists {
		return "key is not exist", false
	}

	if len(entry.Values) == 0 {
		return "list is empty", false
	}

	if !entry.expiration.IsZero() && time.Now().After(entry.expiration) {
		return "list entity is expired", false
	}

	lastIndex := len(entry.Values) - 1
	value := entry.Values[lastIndex]
	entry.Values = entry.Values[:lastIndex] // Remove the last element

	s.Lists[key] = entry

	return value, true
}
