// Package globalprotocolsbgpaddressfamilyipvsixunicastlabelvpnallocationmode code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvsixunicastlabelvpnallocationmode

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpAddressFamilyIPvsixUnicastLabelVpnAllocationMode{}
	_ resource.ResourceWithConfigure = &protocolsBgpAddressFamilyIPvsixUnicastLabelVpnAllocationMode{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvsixUnicastLabelVpnAllocationMode{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvsixUnicastLabelVpnAllocationMode{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvsixUnicastLabelVpnAllocationMode{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvsixUnicastLabelVpnAllocationMode{}
// var _ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvsixUnicastLabelVpnAllocationMode{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvsixUnicastLabelVpnAllocationMode{}