/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedprotocolsfailoverroute code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsfailoverroute

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsFailoverRoute{}
	_ resource.ResourceWithConfigure   = &protocolsFailoverRoute{}
	_ resource.ResourceWithImportState = &protocolsFailoverRoute{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsFailoverRoute{}
// var _ resource.ResourceWithModifyPlan = &protocolsFailoverRoute{}
// var _ resource.ResourceWithUpgradeState = &protocolsFailoverRoute{}
// var _ resource.ResourceWithValidateConfig = &protocolsFailoverRoute{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsFailoverRoute{}
