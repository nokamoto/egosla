package mysql

import "errors"

var (
	// ErrInvalidArgument represents an error which arguments is invalid.
	ErrInvalidArgument = errors.New("invalid argument")
	// ErrUnknown represents an error which is unneccessary to distinguish the cause or just unexpected.
	ErrUnknown = errors.New("unknown")
)
