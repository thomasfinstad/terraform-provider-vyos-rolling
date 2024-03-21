package cruderrors

import (
	"errors"
	"fmt"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// ResourceNotFoundError indicates the client could not find the requested config
type ResourceNotFoundError struct {
	error
	model helpers.VyosTopResourceDataModel
}

func (e ResourceNotFoundError) Error() string {
	return fmt.Sprintf("[%s] resource not found", e.model.GetVyosPath())
}

// Unwrap returns the underlyding error
func Unwrap(err error) error {
	if e, ok := err.(ResourceNotFoundError); ok {
		return errors.Unwrap(e.error)
	}

	return errors.Unwrap(err)
}

// NewResourceNotFoundError formates and returns a new ResourceNotFoundError
func NewResourceNotFoundError(model helpers.VyosTopResourceDataModel, format string, args ...interface{}) error {
	return ResourceNotFoundError{
		error: fmt.Errorf(format, args...),
		model: model,
	}
}

// WrapIntoResourceNotFoundError wraps an error and returns it as a ResourceNotFoundError
func WrapIntoResourceNotFoundError(model helpers.VyosTopResourceDataModel, err error) error {
	return ResourceNotFoundError{
		model: model,
		error: fmt.Errorf("%w", err),
	}
}

// Details returns the underlying details of the error
func (e ResourceNotFoundError) Details() string {
	return e.error.Error()
}
