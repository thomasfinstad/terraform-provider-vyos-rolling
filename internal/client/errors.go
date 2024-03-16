package client

import "fmt"

// APINotFoundError desiered path was not found when reading the API
type APINotFoundError struct {
	message string
	path    string
}

// Error returns human friendly version of error
func (e *APINotFoundError) Error() string {
	return fmt.Sprintf("[api error] '%s': '%s'", e.path, e.message)
}

// MarshalError unable to marshal
type MarshalError struct {
	message    string
	marshalErr error
}

// Error returns human friendly version of error
func (e *MarshalError) Error() string {
	return fmt.Sprintf("[marshal error] %s: %w", e.message, e.marshalErr)
}
