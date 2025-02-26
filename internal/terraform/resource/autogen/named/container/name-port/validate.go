/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedcontainernameport code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedcontainernameport

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (port) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &containerNamePort{}
	_ resource.ResourceWithConfigure   = &containerNamePort{}
	_ resource.ResourceWithImportState = &containerNamePort{}
)

// var _ resource.ResourceWithConfigValidators = &containerNamePort{}
// var _ resource.ResourceWithModifyPlan = &containerNamePort{}
// var _ resource.ResourceWithUpgradeState = &containerNamePort{}
// var _ resource.ResourceWithValidateConfig = &containerNamePort{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &containerNamePort{}
