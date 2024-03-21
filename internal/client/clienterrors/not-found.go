package clienterrors

import (
	"errors"
	"fmt"
)

// NotFoundError indicates the client could not find the requested config
type NotFoundError struct {
	error
}

func (e NotFoundError) Error() string {
	return e.error.Error()
}

// Unwrap returns the underlyding error
func Unwrap(err error) error {
	if e, ok := err.(NotFoundError); ok {
		return errors.Unwrap(e.error)
	}

	return errors.Unwrap(err)
}

// NewNotFoundError formates and returns a new NotFoundError
func NewNotFoundError(format string, args ...interface{}) NotFoundError {
	return NotFoundError{
		error: fmt.Errorf(format, args...),
	}
}

// WrapIntoNotFoundError wraps an error and returns it as a NotFoundError
func WrapIntoNotFoundError(msg string, err error) NotFoundError {
	return NotFoundError{
		error: fmt.Errorf("[%s]: %w", msg, err),
	}
}
