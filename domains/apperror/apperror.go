package apperror

import "errors"

var (
	ErrInternal = errors.New("internal")
	ErrInvalid  = errors.New("invalid")
)
