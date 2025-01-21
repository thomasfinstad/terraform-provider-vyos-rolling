/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfaceswwan code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswwan

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (wwan) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesWwan{}
	_ resource.ResourceWithConfigure   = &interfacesWwan{}
	_ resource.ResourceWithImportState = &interfacesWwan{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesWwan{}
// var _ resource.ResourceWithModifyPlan = &interfacesWwan{}
// var _ resource.ResourceWithUpgradeState = &interfacesWwan{}
// var _ resource.ResourceWithValidateConfig = &interfacesWwan{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesWwan{}
