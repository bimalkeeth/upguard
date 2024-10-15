package constants

import "errors"

var (
	ErrInvalidBatchSize = errors.New("invalid batch size")
	ErrJobCannotBeNil   = errors.New("job cannot be nil")
	ErrTimeDuration     = errors.New("time duration cannot be negative or zero")
)
