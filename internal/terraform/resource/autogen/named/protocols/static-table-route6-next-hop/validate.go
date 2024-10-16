// Package namedprotocolsstatictableroutesixnexthop code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstatictableroutesixnexthop

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsStaticTableRoutesixNextHop{}
	_ resource.ResourceWithConfigure   = &protocolsStaticTableRoutesixNextHop{}
	_ resource.ResourceWithImportState = &protocolsStaticTableRoutesixNextHop{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticTableRoutesixNextHop{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticTableRoutesixNextHop{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticTableRoutesixNextHop{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticTableRoutesixNextHop{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticTableRoutesixNextHop{}
