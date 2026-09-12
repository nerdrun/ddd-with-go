package domain

import "errors"

var (
	ErrProjectFull = errors.New("too many tasks")
)
