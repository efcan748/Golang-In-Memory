package server

import (
	"time"

	"github.com/efcan748/Golang-In-Memory/pkg/store"
)

type Server struct {
	store *store.Store
}

func New(cleanUpTime int) *Server {
	return &Server{
		store: store.New().StartCleanup(time.Duration(cleanUpTime) * time.Minute),
	}
}
