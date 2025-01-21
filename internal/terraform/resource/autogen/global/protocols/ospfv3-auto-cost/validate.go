/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsospfvthreeautocost code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfvthreeautocost

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (auto-cost) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfvthreeAutoCost{}
	_ resource.ResourceWithConfigure   = &protocolsOspfvthreeAutoCost{}
	_ resource.ResourceWithImportState = &protocolsOspfvthreeAutoCost{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfvthreeAutoCost{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfvthreeAutoCost{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfvthreeAutoCost{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfvthreeAutoCost{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfvthreeAutoCost{}
