/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsripdistributelistaccesslist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripdistributelistaccesslist

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsRIPDistributeListAccessList{}
	_ resource.ResourceWithConfigure   = &protocolsRIPDistributeListAccessList{}
	_ resource.ResourceWithImportState = &protocolsRIPDistributeListAccessList{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPDistributeListAccessList{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPDistributeListAccessList{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPDistributeListAccessList{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPDistributeListAccessList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPDistributeListAccessList{}
