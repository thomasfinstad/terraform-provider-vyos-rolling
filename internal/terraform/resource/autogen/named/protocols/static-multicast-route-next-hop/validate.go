// Package namedprotocolsstaticmulticastroutenexthop code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstaticmulticastroutenexthop

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsStaticMulticastRouteNextHop{}
	_ resource.ResourceWithConfigure = &protocolsStaticMulticastRouteNextHop{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticMulticastRouteNextHop{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticMulticastRouteNextHop{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticMulticastRouteNextHop{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticMulticastRouteNextHop{}
// var _ resource.ResourceWithImportState = &protocolsStaticMulticastRouteNextHop{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticMulticastRouteNextHop{}