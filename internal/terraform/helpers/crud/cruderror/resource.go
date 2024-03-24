package cruderrors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// ResourceError indicates the client could not find the requested config
//
// This implements the error interface, as well as the terraform diag.Diagnostics
// interface so it can be used seemlesly as either.
// To create a ResourceError you can use either `NewResourceError`
// to create a new one from scratch, or `WrapIntoResourceError` to wrap
// an existing error and return it as a ResourceError.
type ResourceError struct {
	error
	model helpers.VyosTopResourceDataModel
}

/*
--------------------------------------------
diag.Diagnostics interface implementation
--------------------------------------------
*/

// NewResourceError formats and returns a new ResourceError
//
// msgFormat  : in the format of `fmt.Errorf()`
// args       : values used in msgFormat
func NewResourceError(model helpers.VyosTopResourceDataModel, msgFormat string, args ...interface{}) ResourceError {
	return ResourceError{
		error: fmt.Errorf(msgFormat, args...),
		model: model,
	}
}

// WrapIntoResourceError wraps an error and returns it as a ResourceError
func WrapIntoResourceError(model helpers.VyosTopResourceDataModel, err error) ResourceError {
	return ResourceError{
		model: model,
		error: fmt.Errorf("%w", err),
	}
}

/*
--------------------------------------------
error interface implementation + unwrap
--------------------------------------------
*/

func (e ResourceError) Error() string {
	return fmt.Sprintf("[%s] resource error: %s", e.model.GetVyosPath(), e.error.Error())
}

// Unwrap returns the underlyding error
func (e ResourceError) Unwrap() error {
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
func (e ResourceError) Summary() string {
	return fmt.Sprintf("[%s] resource error", e.model.GetVyosPath())
}

// Detail is a long description for the diagnostic.
//
// This should contain all relevant information about why the diagnostic
// was generated and if applicable, ways to prevent the diagnostic. It
// should generally be written and formatted for human consumption by
// practitioners or provider developers.
func (e ResourceError) Detail() string {
	return e.error.Error()
}

// Severity returns the desired level of feedback for the diagnostic.
func (e ResourceError) Severity() diag.Severity {
	return diag.SeverityError
}

// Equal returns true if the other diagnostic is wholly equivalent.
func (e ResourceError) Equal(other diag.Diagnostic) bool {

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

	if e.error.Error() != o.error.Error() {
		return false
	}

	if strings.Join(e.model.GetVyosPath(), "") != strings.Join(o.model.GetVyosPath(), "") {
		return false
	}

	if e.error.Error() != o.error.Error() {
		return false
	}

	return true
}
