package api

// import (
// 	"github.com/efcan748/Golang-In-Memory/internal/client"
// )

// type Handler struct {
// 	client *client.Client
// }

// func NewHandler(client *client.Client) *Handler {
// 	return &Handler{client: client}
// }

// func (h *Handler) Set(w http.ResponseWriter, r *http.Request) {
// 	var req models.SetRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "invalid request")
// 		return
// 	}

// 	if err := h.client.Set(r.Context(), req.Key, req.Value, req.TTL); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
// }

// func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
// 	key := r.URL.Query().Get("key")
// 	if key == "" {
// 		respondWithError(w, http.StatusBadRequest, "key is required")
// 		return
// 	}

// 	value, err := h.client.Get(r.Context(), key)
// 	if err != nil {
// 		if err == client.ErrKeyNotFound {
// 			respondWithError(w, http.StatusNotFound, "key not found")
// 			return
// 		}
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, models.GetResponse{Value: value})
// }

// func (h *Handler) PushList(w http.ResponseWriter, r *http.Request) {
// 	var req models.SetRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "invalid request")
// 		return
// 	}

// 	if err := h.client.PushList(r.Context(), req.Key, req.TTL, req.Value); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
// }

// func (h *Handler) PopList(w http.ResponseWriter, r *http.Request) {
// 	key := r.URL.Query().Get("key")
// 	if key == "" {
// 		respondWithError(w, http.StatusBadRequest, "key is required")
// 		return
// 	}

// 	value, err := h.client.PopList(r.Context(), key)
// 	if err != nil {
// 		if err == client.ErrKeyNotFound {
// 			respondWithError(w, http.StatusNotFound, "key not found")
// 			return
// 		}
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, models.GetResponse{Value: value})
// }

// // Similar handlers for Delete, LPush, LPop...
