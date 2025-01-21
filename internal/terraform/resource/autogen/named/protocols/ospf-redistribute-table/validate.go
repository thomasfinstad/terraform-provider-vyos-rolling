/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsospfredistributetable code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfredistributetable

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (table) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfRedistributeTable{}
	_ resource.ResourceWithConfigure   = &protocolsOspfRedistributeTable{}
	_ resource.ResourceWithImportState = &protocolsOspfRedistributeTable{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfRedistributeTable{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfRedistributeTable{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfRedistributeTable{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfRedistributeTable{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfRedistributeTable{}
