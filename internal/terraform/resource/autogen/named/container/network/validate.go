/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedcontainernetwork code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedcontainernetwork

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (network) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &containerNetwork{}
	_ resource.ResourceWithConfigure   = &containerNetwork{}
	_ resource.ResourceWithImportState = &containerNetwork{}
)

// var _ resource.ResourceWithConfigValidators = &containerNetwork{}
// var _ resource.ResourceWithModifyPlan = &containerNetwork{}
// var _ resource.ResourceWithUpgradeState = &containerNetwork{}
// var _ resource.ResourceWithValidateConfig = &containerNetwork{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &containerNetwork{}
