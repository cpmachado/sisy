package todotxt

import "errors"

var (
	ErrInvalidPriority     = errors.New("invalid priority")
	ErrMissingCreationDate = errors.New("missing creation date")
)
