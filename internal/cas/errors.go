package cas

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrHashMismatch = errors.New("hash mismatch")
)
