package logic

import "errors"

var (
	ErrAssertionFailed     = errors.New("assertion failed")
	ErrQuerryLengthInvalid = errors.New("querry length invalid")
)
