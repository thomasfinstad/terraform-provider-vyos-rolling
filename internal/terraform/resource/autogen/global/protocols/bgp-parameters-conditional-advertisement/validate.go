/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsbgpparametersconditionaladvertisement code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpparametersconditionaladvertisement

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpParametersConditionalAdvertisement{}
	_ resource.ResourceWithConfigure   = &protocolsBgpParametersConditionalAdvertisement{}
	_ resource.ResourceWithImportState = &protocolsBgpParametersConditionalAdvertisement{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpParametersConditionalAdvertisement{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpParametersConditionalAdvertisement{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpParametersConditionalAdvertisement{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpParametersConditionalAdvertisement{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpParametersConditionalAdvertisement{}
