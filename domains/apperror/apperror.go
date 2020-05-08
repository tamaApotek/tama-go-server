package apperror

import "errors"

var (
	ErrInternal = errors.New("internal")
	ErrInvalid  = errors.New("invalid")
	ErrNotFound = errors.New("not-found")
)
