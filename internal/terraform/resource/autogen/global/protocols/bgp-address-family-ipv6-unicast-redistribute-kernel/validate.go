// Package globalprotocolsbgpaddressfamilyipvsixunicastredistributekernel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvsixunicastredistributekernel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel{}
	_ resource.ResourceWithConfigure   = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel{}
	_ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel{}
