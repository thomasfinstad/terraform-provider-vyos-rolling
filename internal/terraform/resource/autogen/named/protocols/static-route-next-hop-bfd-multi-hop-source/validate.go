// Package namedprotocolsstaticroutenexthopbfdmultihopsource code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstaticroutenexthopbfdmultihopsource

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsStaticRouteNextHopBfdMultiHopSource{}
	_ resource.ResourceWithConfigure   = &protocolsStaticRouteNextHopBfdMultiHopSource{}
	_ resource.ResourceWithImportState = &protocolsStaticRouteNextHopBfdMultiHopSource{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticRouteNextHopBfdMultiHopSource{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticRouteNextHopBfdMultiHopSource{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticRouteNextHopBfdMultiHopSource{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticRouteNextHopBfdMultiHopSource{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticRouteNextHopBfdMultiHopSource{}
