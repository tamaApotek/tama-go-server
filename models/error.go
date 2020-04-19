package models

import (
	"errors"
)

var (
	// ErrInternal Internal server error
	ErrInternal = errors.New("internal")
	// ErrEmpty Requested data is null
	ErrEmpty = errors.New("empty")
	// ErrInvalid Invalid data payload
	ErrInvalid = errors.New("invalid")
	// ErrConflict Data already exists
	ErrConflict = errors.New("conflict")
	// ErrNotFound 404
	ErrNotFound = errors.New("not_found")
)
