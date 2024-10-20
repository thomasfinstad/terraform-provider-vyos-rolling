/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsospfldpsync code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfldpsync

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfLdpSync{}
	_ resource.ResourceWithConfigure   = &protocolsOspfLdpSync{}
	_ resource.ResourceWithImportState = &protocolsOspfLdpSync{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfLdpSync{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfLdpSync{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfLdpSync{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfLdpSync{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfLdpSync{}
