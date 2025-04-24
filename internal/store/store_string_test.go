package store

import (
	"testing"
	"time"
)

func TestSetGetDelete(t *testing.T) {
	s := New()

	// Set
	msg, ok := s.Set("key1", "value1", 2*time.Second)
	if !ok {
		t.Fatalf("Set failed: %s", msg)
	}

	// Get
	val, ok := s.Get("key1")
	if !ok || val != "value1" {
		t.Errorf("Get failed: expected value1, got %s", val)
	}

	// Delete
	if !s.Delete("key1") {
		t.Errorf("Delete failed on existing key")
	}

	// Get after delete
	_, ok = s.Get("key1")
	if ok {
		t.Errorf("Expected Get to fail on deleted key")
	}
}

func TestSetExistingKey(t *testing.T) {
	s := New()

	s.Set("key1", "value1", 0)
	msg, ok := s.Set("key1", "newvalue", 0)
	if ok {
		t.Errorf("Expected Set to fail on existing key")
	}
	if msg == "" {
		t.Errorf("Expected error message when setting existing key")
	}
}

func TestUpdate(t *testing.T) {
	s := New()

	// Update non-existent key
	msg, ok := s.Update("keyX", "valueX", 0)
	if ok {
		t.Errorf("Expected Update to fail on non-existent key")
	}
	if msg == "" {
		t.Errorf("Expected message on failed update")
	}

	// Set and Update
	s.Set("key2", "value2", 0)
	msg, ok = s.Update("key2", "updated", 0)
	if !ok {
		t.Errorf("Update failed: %s", msg)
	}

	val, ok := s.Get("key2")
	if !ok || val != "updated" {
		t.Errorf("Expected updated value2, got %s", val)
	}
}

func TestExpiration(t *testing.T) {
	s := New()

	s.Set("tempKey", "tempVal", 1*time.Second)

	time.Sleep(2 * time.Second) // Let it expire

	_, ok := s.Get("tempKey")
	if ok {
		t.Errorf("Expected key to be expired")
	}

	msg, ok := s.Update("tempKey", "newVal", 0)
	if ok {
		t.Errorf("Expected Update to fail on expired key")
	}
	if msg != "entity is expired" {
		t.Errorf("Expected expiration message, got %s", msg)
	}
}
