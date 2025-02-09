/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfacesdummy code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesdummy

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (dummy) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesDummy{}
	_ resource.ResourceWithConfigure   = &interfacesDummy{}
	_ resource.ResourceWithImportState = &interfacesDummy{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesDummy{}
// var _ resource.ResourceWithModifyPlan = &interfacesDummy{}
// var _ resource.ResourceWithUpgradeState = &interfacesDummy{}
// var _ resource.ResourceWithValidateConfig = &interfacesDummy{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesDummy{}
