// Package globalprotocolsbabeldistributelistipvfourprefixlist code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbabeldistributelistipvfourprefixlist

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBabelDistributeListIPvfourPrefixList{}
	_ resource.ResourceWithConfigure = &protocolsBabelDistributeListIPvfourPrefixList{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBabelDistributeListIPvfourPrefixList{}
// var _ resource.ResourceWithModifyPlan = &protocolsBabelDistributeListIPvfourPrefixList{}
// var _ resource.ResourceWithUpgradeState = &protocolsBabelDistributeListIPvfourPrefixList{}
// var _ resource.ResourceWithValidateConfig = &protocolsBabelDistributeListIPvfourPrefixList{}
// var _ resource.ResourceWithImportState = &protocolsBabelDistributeListIPvfourPrefixList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBabelDistributeListIPvfourPrefixList{}
