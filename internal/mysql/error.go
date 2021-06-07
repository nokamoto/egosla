package mysql

import "errors"

var (
	// ErrInvalidArgument represents an error which arguments is invalid.
	ErrInvalidArgument = errors.New("invalid argument")
	// ErrNotFound represents specified records not found.
	ErrNotFound = errors.New("not found")
	// ErrUnknown represents an error which is unneccessary to distinguish the cause or just unexpected.
	ErrUnknown = errors.New("unknown")
)
