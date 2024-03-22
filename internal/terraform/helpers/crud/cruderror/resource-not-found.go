package cruderrors

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// ResourceNotFoundError indicates the client could not find the requested config
//
// This implements the error interface, as well as the terraform diag.Diagnostics
// interface so it can be used seemlesly as either.
// To create a ResourceNotFoundError you can use either `NewResourceNotFoundError`
// to create a new one from scratch, or `WrapIntoResourceNotFoundError` to wrap
// an existing error and return it as a ResourceNotFoundError.
type ResourceNotFoundError struct {
	error
	model helpers.VyosTopResourceDataModel
}

/*
--------------------------------------------
diag.Diagnostics interface implementation
--------------------------------------------
*/

// NewResourceNotFoundError formates and returns a new ResourceNotFoundError
func NewResourceNotFoundError(model helpers.VyosTopResourceDataModel, format string, args ...interface{}) ResourceNotFoundError {
	return ResourceNotFoundError{
		error: fmt.Errorf(format, args...),
		model: model,
	}
}

// WrapIntoResourceNotFoundError wraps an error and returns it as a ResourceNotFoundError
func WrapIntoResourceNotFoundError(model helpers.VyosTopResourceDataModel, err error) ResourceNotFoundError {
	return ResourceNotFoundError{
		model: model,
		error: fmt.Errorf("%w", err),
	}
}

/*
--------------------------------------------
error interface implementation + unwrap
--------------------------------------------
*/

func (e ResourceNotFoundError) Error() string {
	return fmt.Sprintf("[%s] resource error: %s", e.model.GetVyosPath(), e.error.Error())
}

// Unwrap returns the underlyding error
func (e ResourceNotFoundError) Unwrap() error {
	return errors.Unwrap(e.error)
}

/*
--------------------------------------------
diag.Diagnostics interface implementation
--------------------------------------------
*/

// Summary is a short description for the diagnostic.
//
// Typically this is implemented as a title, such as "Invalid Resource Name",
// or single line sentence.
func (e ResourceNotFoundError) Summary() string {
	return e.Error()
}

// Detail is a long description for the diagnostic.
//
// This should contain all relevant information about why the diagnostic
// was generated and if applicable, ways to prevent the diagnostic. It
// should generally be written and formatted for human consumption by
// practitioners or provider developers.
func (e ResourceNotFoundError) Detail() string {
	return e.error.Error()
}

// Severity returns the desired level of feedback for the diagnostic.
func (e ResourceNotFoundError) Severity() diag.Severity {
	return diag.SeverityError
}

// Equal returns true if the other diagnostic is wholly equivalent.
func (e ResourceNotFoundError) Equal(other diag.Diagnostic) bool {

	o, ok := other.(ResourceNotFoundError)

	if !ok {
		return false
	}

	if e.error == nil {
		return o.error == nil
	} else if o.error == nil {
		return false
	}

	if e.model == nil {
		return o.model == nil
	} else if o.model == nil {
		return false
	}

	if e.Detail() != o.Detail() {
		return false
	}

	if e.Summary() != o.Summary() {
		return false
	}

	if e.error.Error() != o.error.Error() {
		return false
	}

	return true
}
