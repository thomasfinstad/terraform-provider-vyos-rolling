// Package globalprotocolsbgpaddressfamilyipvfourunicastlabelvpnallocationmode code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvfourunicastlabelvpnallocationmode

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode{}
	_ resource.ResourceWithConfigure   = &protocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode{}
	_ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvfourUnicastLabelVpnAllocationMode{}
