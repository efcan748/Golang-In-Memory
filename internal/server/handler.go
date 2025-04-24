package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/efcan748/Golang-In-Memory/internal/store"
	"github.com/efcan748/Golang-In-Memory/pkg/models"
)

type Handler struct {
	store *store.Store
}

func New() *Handler {
	return &Handler{
		store: store.New().StartCleanup(1 * time.Minute),
	}
}

func (c *Handler) Get(w http.ResponseWriter, r *http.Request) {
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

func (c *Handler) Set(w http.ResponseWriter, r *http.Request) {
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

func (c *Handler) Delete(w http.ResponseWriter, r *http.Request, id string) {

	success := c.store.Delete(id)
	if !success {
		respondWithError(w, http.StatusNotFound, "key not found")
	} else {
		respondWithJSON(w, http.StatusOK, map[string]string{"status": "successfully removed."})

	}

}

func (c *Handler) Update(w http.ResponseWriter, r *http.Request, id string) {
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

func (c *Handler) PushList(w http.ResponseWriter, r *http.Request) {
	var req models.SetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	ttlMilisecond := time.Duration(req.TTL) * time.Millisecond
	msg, success := c.store.LPush(req.Key, ttlMilisecond, req.Value)

	if !success {
		respondWithError(w, http.StatusBadRequest, msg)
	} else {
		respondWithJSON(w, http.StatusOK, map[string]string{"status": "successfully stored."})
	}

}

func (c *Handler) PopList(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		respondWithError(w, http.StatusBadRequest, "key is required")
		return
	}

	value, exists := c.store.Pop(key)
	if !exists {
		respondWithError(w, http.StatusNotFound, value)
		return
	}

	fmt.Println(value, "==========")

	respondWithJSON(w, http.StatusOK, models.GetResponse{Value: value})
}
