/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedsystemsyslogremotefacility code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemsyslogremotefacility

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (facility) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemSyslogRemoteFacility{}
	_ resource.ResourceWithConfigure   = &systemSyslogRemoteFacility{}
	_ resource.ResourceWithImportState = &systemSyslogRemoteFacility{}
)

// var _ resource.ResourceWithConfigValidators = &systemSyslogRemoteFacility{}
// var _ resource.ResourceWithModifyPlan = &systemSyslogRemoteFacility{}
// var _ resource.ResourceWithUpgradeState = &systemSyslogRemoteFacility{}
// var _ resource.ResourceWithValidateConfig = &systemSyslogRemoteFacility{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemSyslogRemoteFacility{}
