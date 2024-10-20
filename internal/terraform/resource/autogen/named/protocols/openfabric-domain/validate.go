/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedprotocolsopenfabricdomain code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsopenfabricdomain

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOpenfabricDomain{}
	_ resource.ResourceWithConfigure   = &protocolsOpenfabricDomain{}
	_ resource.ResourceWithImportState = &protocolsOpenfabricDomain{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOpenfabricDomain{}
// var _ resource.ResourceWithModifyPlan = &protocolsOpenfabricDomain{}
// var _ resource.ResourceWithUpgradeState = &protocolsOpenfabricDomain{}
// var _ resource.ResourceWithValidateConfig = &protocolsOpenfabricDomain{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOpenfabricDomain{}
