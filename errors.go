package postal

import "errors"

var (
	ErrInvalidCode   = errors.New("supplied code is invalid")
	ErrStateNotFound = errors.New("state not found by supplied code")
)
