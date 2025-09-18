package jsonbagger

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrNestingOverflow = errors.New("nesting level overflow")
)
