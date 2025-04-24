package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/efcan748/Golang-In-Memory/pkg/models"
)

func setupTestServer() http.Handler {
	s := New(1)
	return NewRouter(s)
}

func TestStringSetGetDelete(t *testing.T) {
	handler := setupTestServer()

	// 1. Test Set
	body := models.SetRequest{
		Key:   "foo",
		Value: "bar",
		TTL:   1000,
	}
	data, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/string", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK from Set, got %d", rec.Code)
	}

	// 2. Test Get
	req = httptest.NewRequest(http.MethodGet, "/string?key=foo", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK from Get, got %d", rec.Code)
	}

	var res models.GetResponse
	bodyBytes, _ := io.ReadAll(rec.Body)
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		t.Fatalf("Failed to parse JSON response: %v", err)
	}
	if res.Value != "bar" {
		t.Errorf("Expected value 'bar', got '%s'", res.Value)
	}

	// 3. Test Delete
	req = httptest.NewRequest(http.MethodDelete, "/string/foo", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK from Delete, got %d", rec.Code)
	}

	// 4. Test Get after Delete (should 404)
	req = httptest.NewRequest(http.MethodGet, "/string?key=foo", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected 404 Not Found after delete, got %d", rec.Code)
	}
}

func TestUpdateEndpoint(t *testing.T) {
	handler := setupTestServer()

	// Set first
	setReq := models.SetRequest{
		Key:   "upkey",
		Value: "original",
		TTL:   0,
	}
	data, _ := json.Marshal(setReq)
	req := httptest.NewRequest(http.MethodPost, "/string", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	// Now Update
	upReq := models.UpdateRequest{
		Value: "updated",
		TTL:   0,
	}
	upData, _ := json.Marshal(upReq)
	req = httptest.NewRequest(http.MethodPut, "/string/upkey", bytes.NewReader(upData))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK from Update, got %d", rec.Code)
	}

	// Confirm update worked
	req = httptest.NewRequest(http.MethodGet, "/string?key=upkey", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	var res models.GetResponse
	bodyBytes, _ := io.ReadAll(rec.Body)
	json.Unmarshal(bodyBytes, &res)

	if res.Value != "updated" {
		t.Errorf("Expected updated value 'updated', got '%s'", res.Value)
	}
}

func TestSetTTLExpiration(t *testing.T) {
	handler := setupTestServer()

	body := models.SetRequest{
		Key:   "temp",
		Value: "data",
		TTL:   500, // ms
	}
	data, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/string", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	// Wait for expiration
	time.Sleep(600 * time.Millisecond)

	// Try to Get after expiration
	req = httptest.NewRequest(http.MethodGet, "/string?key=temp", nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected 404 after TTL expiration, got %d", rec.Code)
	}
}
