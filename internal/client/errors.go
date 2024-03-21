package client

import (
	"fmt"
)

// MarshalError unable to marshal
type MarshalError struct {
	message    string
	marshalErr error
}

// Error returns human friendly version of error
func (e *MarshalError) Error() string {
	return fmt.Sprintf("[marshal error] %s: %s", e.message, e.marshalErr)
}
