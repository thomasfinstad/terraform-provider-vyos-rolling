// Package globalprotocolsbabeldistributelistipvsixaccesslist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbabeldistributelistipvsixaccesslist

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBabelDistributeListIPvsixAccessList{}
	_ resource.ResourceWithConfigure   = &protocolsBabelDistributeListIPvsixAccessList{}
	_ resource.ResourceWithImportState = &protocolsBabelDistributeListIPvsixAccessList{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBabelDistributeListIPvsixAccessList{}
// var _ resource.ResourceWithModifyPlan = &protocolsBabelDistributeListIPvsixAccessList{}
// var _ resource.ResourceWithUpgradeState = &protocolsBabelDistributeListIPvsixAccessList{}
// var _ resource.ResourceWithValidateConfig = &protocolsBabelDistributeListIPvsixAccessList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBabelDistributeListIPvsixAccessList{}
