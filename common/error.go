package common

import "errors"

var (
	ErrInternal    = errors.New("internal")
	ErrInvalid     = errors.New("invalid")
	ErrInvalidBody = errors.New("invalid-body")
	ErrNotFound    = errors.New("not-found")
)
