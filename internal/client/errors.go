package client

import "errors"

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrInvalidType = errors.New("invalid type for operation")
	ErrEmptyList   = errors.New("list is empty")
	// Add other error types as needed
)
