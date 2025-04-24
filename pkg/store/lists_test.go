package store_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/efcan748/Golang-In-Memory/pkg/store"
)

func TestLPush(t *testing.T) {
	s := store.New()

	tests := []struct {
		key      string
		ttl      time.Duration
		values   []string
		expected []string
	}{
		{
			key:      "test1",
			ttl:      0,
			values:   []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			key:      "test2",
			ttl:      time.Second * 5,
			values:   []string{"x", "y"},
			expected: []string{"x", "y"},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("LPush %s", tt.key), func(t *testing.T) {
			_, ok := s.LPush(tt.key, tt.ttl, tt.values...)
			if !ok {
				t.Fatalf("Expected true but got false")
			}

			entry := s.Lists[tt.key]
			if len(entry.Values) != len(tt.expected) {
				t.Errorf("Expected length %d but got %d", len(tt.expected), len(entry.Values))
			}

			for i, v := range entry.Values {
				if v != tt.expected[i] {
					t.Errorf("Expected value %s at index %d but got %s", tt.expected[i], i, v)
				}
			}
		})
	}
}

func TestPop(t *testing.T) {
	s := store.New()

	s.LPush("popTest", 0, "first", "second")

	value, ok := s.Pop("popTest")
	if !ok || value != "second" {
		t.Errorf("Expected 'second', got '%s'", value)
	}

	value, ok = s.Pop("popTest")
	if !ok || value != "first" {
		t.Errorf("Expected 'first', got '%s'", value)
	}

	_, ok = s.Pop("popTest")
	if ok {
		t.Error("Expected false for empty list pop, but got true")
	}
}
