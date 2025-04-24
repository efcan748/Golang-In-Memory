package models

import "time"

// Request/Response models for API endpoints

type SetRequest struct {
	Key   string        `json:"key"`
	Value string        `json:"value"`
	TTL   time.Duration `json:"ttl"` // Duration in nanoseconds
}

type GetResponse struct {
	Value string `json:"value"`
}

type UpdateRequest struct {
	Value string        `json:"value"`
	TTL   time.Duration `json:"ttl"` // Duration in nanoseconds`
}

type ListOperationRequest struct {
	Key    string   `json:"key"`
	Values []string `json:"values,omitempty"` // Used for LPush
}

type ListResponse struct {
	Value string `json:"value"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
