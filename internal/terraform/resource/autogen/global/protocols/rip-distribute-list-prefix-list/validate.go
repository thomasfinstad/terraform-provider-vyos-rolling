// Package globalprotocolsripdistributelistprefixlist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripdistributelistprefixlist

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsRIPDistributeListPrefixList{}
	_ resource.ResourceWithConfigure = &protocolsRIPDistributeListPrefixList{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPDistributeListPrefixList{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPDistributeListPrefixList{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPDistributeListPrefixList{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPDistributeListPrefixList{}
// var _ resource.ResourceWithImportState = &protocolsRIPDistributeListPrefixList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPDistributeListPrefixList{}
