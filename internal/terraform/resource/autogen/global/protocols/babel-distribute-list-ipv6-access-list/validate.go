// Package globalprotocolsbabeldistributelistipvsixaccesslist code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbabeldistributelistipvsixaccesslist

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBabelDistributeListIPvsixAccessList{}
	_ resource.ResourceWithConfigure = &protocolsBabelDistributeListIPvsixAccessList{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBabelDistributeListIPvsixAccessList{}
// var _ resource.ResourceWithModifyPlan = &protocolsBabelDistributeListIPvsixAccessList{}
// var _ resource.ResourceWithUpgradeState = &protocolsBabelDistributeListIPvsixAccessList{}
// var _ resource.ResourceWithValidateConfig = &protocolsBabelDistributeListIPvsixAccessList{}
// var _ resource.ResourceWithImportState = &protocolsBabelDistributeListIPvsixAccessList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBabelDistributeListIPvsixAccessList{}
