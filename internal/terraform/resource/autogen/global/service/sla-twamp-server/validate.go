/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalserviceslatwampserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceslatwampserver

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (twamp-server) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceSLATwampServer{}
	_ resource.ResourceWithConfigure   = &serviceSLATwampServer{}
	_ resource.ResourceWithImportState = &serviceSLATwampServer{}
)

// var _ resource.ResourceWithConfigValidators = &serviceSLATwampServer{}
// var _ resource.ResourceWithModifyPlan = &serviceSLATwampServer{}
// var _ resource.ResourceWithUpgradeState = &serviceSLATwampServer{}
// var _ resource.ResourceWithValidateConfig = &serviceSLATwampServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceSLATwampServer{}
