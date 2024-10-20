/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedprotocolsospfneighbor code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfneighbor

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfNeighbor{}
	_ resource.ResourceWithConfigure   = &protocolsOspfNeighbor{}
	_ resource.ResourceWithImportState = &protocolsOspfNeighbor{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfNeighbor{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfNeighbor{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfNeighbor{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfNeighbor{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfNeighbor{}
