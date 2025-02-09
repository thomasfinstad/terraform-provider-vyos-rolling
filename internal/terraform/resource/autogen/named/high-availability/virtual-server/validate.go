/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedhighavailabilityvirtualserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedhighavailabilityvirtualserver

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (virtual-server) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &highAvailabilityVirtualServer{}
	_ resource.ResourceWithConfigure   = &highAvailabilityVirtualServer{}
	_ resource.ResourceWithImportState = &highAvailabilityVirtualServer{}
)

// var _ resource.ResourceWithConfigValidators = &highAvailabilityVirtualServer{}
// var _ resource.ResourceWithModifyPlan = &highAvailabilityVirtualServer{}
// var _ resource.ResourceWithUpgradeState = &highAvailabilityVirtualServer{}
// var _ resource.ResourceWithValidateConfig = &highAvailabilityVirtualServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &highAvailabilityVirtualServer{}
