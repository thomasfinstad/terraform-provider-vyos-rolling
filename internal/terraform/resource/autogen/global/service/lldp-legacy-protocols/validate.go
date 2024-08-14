// Package globalservicelldplegacyprotocols code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicelldplegacyprotocols

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceLldpLegacyProtocols{}
	_ resource.ResourceWithConfigure = &serviceLldpLegacyProtocols{}
)

// var _ resource.ResourceWithConfigValidators = &serviceLldpLegacyProtocols{}
// var _ resource.ResourceWithModifyPlan = &serviceLldpLegacyProtocols{}
// var _ resource.ResourceWithUpgradeState = &serviceLldpLegacyProtocols{}
// var _ resource.ResourceWithValidateConfig = &serviceLldpLegacyProtocols{}
// var _ resource.ResourceWithImportState = &serviceLldpLegacyProtocols{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceLldpLegacyProtocols{}