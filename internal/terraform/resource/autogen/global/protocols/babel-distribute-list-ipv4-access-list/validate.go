// Package globalprotocolsbabeldistributelistipvfouraccesslist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbabeldistributelistipvfouraccesslist

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBabelDistributeListIPvfourAccessList{}
	_ resource.ResourceWithConfigure   = &protocolsBabelDistributeListIPvfourAccessList{}
	_ resource.ResourceWithImportState = &protocolsBabelDistributeListIPvfourAccessList{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBabelDistributeListIPvfourAccessList{}
// var _ resource.ResourceWithModifyPlan = &protocolsBabelDistributeListIPvfourAccessList{}
// var _ resource.ResourceWithUpgradeState = &protocolsBabelDistributeListIPvfourAccessList{}
// var _ resource.ResourceWithValidateConfig = &protocolsBabelDistributeListIPvfourAccessList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBabelDistributeListIPvfourAccessList{}
