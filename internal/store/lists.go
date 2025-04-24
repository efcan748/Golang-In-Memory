package store

import (
	"fmt"
	"time"
)

func (s *Store) LPush(key string, ttl time.Duration, values ...string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Println("============L PUSH")

	expiration := time.Time{}
	if ttl > 0 {
		expiration = time.Now().Add(ttl)
	}

	entry := s.lists[key]
	fmt.Println(entry.values, entry.expiration, key)

	s.lists[key] = listEntry{
		values:     append(values, entry.values...),
		expiration: expiration,
	}
	fmt.Println(s.lists[key].values, s.lists[key].expiration, key, "Updated")

}

func (s *Store) Pop(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Println("=============POP")
	entry, exists := s.lists[key]

	fmt.Println(entry, "----------------", exists)
	if !exists {
		return "Key Is not exist", false
	}

	if len(entry.values) == 0 {
		return "Empty list", true
	}

	if !entry.expiration.IsZero() && time.Now().After(entry.expiration) {
		return "List Entity is expired", false
	}

	lastIndex := len(entry.values) - 1
	value := entry.values[lastIndex]
	entry.values = entry.values[:lastIndex] // Remove the last element

	fmt.Println(entry.values, "=========UpDated!!!")
	s.lists[key] = entry

	return value, true
}
