// Package namedprotocolsstaticroutesixnexthop code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstaticroutesixnexthop

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsStaticRoutesixNextHop{}
	_ resource.ResourceWithConfigure = &protocolsStaticRoutesixNextHop{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticRoutesixNextHop{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticRoutesixNextHop{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticRoutesixNextHop{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticRoutesixNextHop{}
// var _ resource.ResourceWithImportState = &protocolsStaticRoutesixNextHop{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticRoutesixNextHop{}
