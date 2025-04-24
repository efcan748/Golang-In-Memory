package server

import (
	"net/http"
	"strings"
)

func NewRouter(h *Server) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/string", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.Set(w, r)
		} else if r.Method == http.MethodGet {
			h.Get(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/string/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 {
			http.Error(w, "Invalid path", http.StatusBadRequest)
			return
		}
		id := parts[2]

		switch r.Method {
		case http.MethodDelete:
			h.Delete(w, r, id)
		case http.MethodPut:
			h.Update(w, r, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		// 	fmt.Println("-----------List Method----------")

		if r.Method == http.MethodPost {
			h.PushList(w, r)
		} else if r.Method == http.MethodGet {
			h.PopList(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
