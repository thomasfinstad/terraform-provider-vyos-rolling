/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsospfrefresh code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfrefresh

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfRefresh{}
	_ resource.ResourceWithConfigure   = &protocolsOspfRefresh{}
	_ resource.ResourceWithImportState = &protocolsOspfRefresh{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfRefresh{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfRefresh{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfRefresh{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfRefresh{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfRefresh{}
