package postal

import "errors"

var (
	// ErrInvalidCode defines an error when the supplied code is not valid.
	ErrInvalidCode = errors.New("supplied code is invalid")

	// ErrStateNotFound defines an error when the supplied code did not match any of the US states.
	ErrStateNotFound = errors.New("state not found by supplied code")
)
