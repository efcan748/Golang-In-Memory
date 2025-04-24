package server

import (
	"time"

	"github.com/efcan748/Golang-In-Memory/internal/store"
)

type Server struct {
	store *store.Store
}

func New() *Server {
	return &Server{
		store: store.New().StartCleanup(1 * time.Minute),
	}
}
