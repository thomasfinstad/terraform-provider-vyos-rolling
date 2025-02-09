/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedsystemsysloguser code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemsysloguser

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (user) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemSyslogUser{}
	_ resource.ResourceWithConfigure   = &systemSyslogUser{}
	_ resource.ResourceWithImportState = &systemSyslogUser{}
)

// var _ resource.ResourceWithConfigValidators = &systemSyslogUser{}
// var _ resource.ResourceWithModifyPlan = &systemSyslogUser{}
// var _ resource.ResourceWithUpgradeState = &systemSyslogUser{}
// var _ resource.ResourceWithValidateConfig = &systemSyslogUser{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemSyslogUser{}
