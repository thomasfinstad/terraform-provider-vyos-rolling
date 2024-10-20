/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsripredistributeospf code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripredistributeospf

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsRIPRedistributeOspf{}
	_ resource.ResourceWithConfigure   = &protocolsRIPRedistributeOspf{}
	_ resource.ResourceWithImportState = &protocolsRIPRedistributeOspf{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPRedistributeOspf{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPRedistributeOspf{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPRedistributeOspf{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPRedistributeOspf{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPRedistributeOspf{}
