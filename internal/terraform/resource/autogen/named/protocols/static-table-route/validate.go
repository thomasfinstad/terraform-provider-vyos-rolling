/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedprotocolsstatictableroute code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstatictableroute

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsStaticTableRoute{}
	_ resource.ResourceWithConfigure   = &protocolsStaticTableRoute{}
	_ resource.ResourceWithImportState = &protocolsStaticTableRoute{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticTableRoute{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticTableRoute{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticTableRoute{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticTableRoute{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticTableRoute{}
