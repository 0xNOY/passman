package models

import (
	"errors"
)

var (
	ErrIDExhaustion  = errors.New("id exhaustion")
	ErrNotExists     = errors.New("it does not exist")
	ErrInternalError = errors.New("internal error")
	ErrAlreadyExists = errors.New("it already exists")
)
