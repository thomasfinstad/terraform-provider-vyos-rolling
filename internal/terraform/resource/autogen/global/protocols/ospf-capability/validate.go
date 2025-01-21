/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsospfcapability code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfcapability

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (capability) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfCapability{}
	_ resource.ResourceWithConfigure   = &protocolsOspfCapability{}
	_ resource.ResourceWithImportState = &protocolsOspfCapability{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfCapability{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfCapability{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfCapability{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfCapability{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfCapability{}
