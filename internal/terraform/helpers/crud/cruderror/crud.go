package cruderrors

import "github.com/hashicorp/terraform-plugin-framework/diag"

// CrudError indicates an error during CRUD operations
// this interface denotes compability as both an error
// and a terraform diag.Diagnostic
type CrudError interface {
	error
	diag.Diagnostic
}
