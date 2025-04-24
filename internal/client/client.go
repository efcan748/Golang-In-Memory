package client

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/efcan748/Golang-In-Memory/internal/store"
	"github.com/efcan748/Golang-In-Memory/pkg/models"
)

type Client struct {
	store *store.Store
}

func New() *Client {
	return &Client{
		store: store.New().StartCleanup(1 * time.Minute),
	}
}

func (c *Client) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		respondWithError(w, http.StatusBadRequest, "key is required")
		return
	}

	value, exists := c.store.Get(key)
	if !exists {
		respondWithError(w, http.StatusNotFound, "key not found")
		return
	}

	respondWithJSON(w, http.StatusOK, models.GetResponse{Value: value})
}

func (c *Client) Set(w http.ResponseWriter, r *http.Request) {
	var req models.SetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	ttlMilisecond := time.Duration(req.TTL) * time.Millisecond
	msg, success := c.store.Set(req.Key, req.Value, ttlMilisecond)

	if !success {
		respondWithError(w, http.StatusBadRequest, msg)
	} else {
		respondWithJSON(w, http.StatusOK, map[string]string{"status": "successfully stored."})
	}
}

func (c *Client) Delete(w http.ResponseWriter, r *http.Request, id string) {

	success := c.store.Delete(id)
	if !success {
		respondWithError(w, http.StatusNotFound, "key not found")
	} else {
		respondWithJSON(w, http.StatusOK, map[string]string{"status": "successfully removed."})

	}

}

func (c *Client) Update(w http.ResponseWriter, r *http.Request, id string) {
	var req models.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	ttlMilisecond := time.Duration(req.TTL) * time.Millisecond
	msg, success := c.store.Update(id, req.Value, ttlMilisecond)

	if !success {
		respondWithError(w, http.StatusBadRequest, msg)
	} else {
		respondWithJSON(w, http.StatusOK, map[string]string{"status": "successfully updated."})
	}

}

// func (c *Client) PushList(ctx context.Context, key string, ttl time.Duration, value ...string) error {
// 	ttlMilisecond := time.Duration(ttl) * time.Millisecond
// 	c.store.LPush(key, ttlMilisecond, value...)

// 	return nil
// }

// func (c *Client) PopList(ctx context.Context, key string) (string, error) {
// 	value, exists := c.store.Pop(key)
// 	if !exists {
// 		return "", ErrKeyNotFound
// 	}
// 	return value, nil
// }

// Other client methods...

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
