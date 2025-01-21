/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedqospolicynetworkemulator code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicynetworkemulator

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (network-emulator) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &qosPolicyNetworkEmulator{}
	_ resource.ResourceWithConfigure   = &qosPolicyNetworkEmulator{}
	_ resource.ResourceWithImportState = &qosPolicyNetworkEmulator{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyNetworkEmulator{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyNetworkEmulator{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyNetworkEmulator{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyNetworkEmulator{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyNetworkEmulator{}
