/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsripng code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripng

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsRIPng{}
	_ resource.ResourceWithConfigure   = &protocolsRIPng{}
	_ resource.ResourceWithImportState = &protocolsRIPng{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPng{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPng{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPng{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPng{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPng{}
