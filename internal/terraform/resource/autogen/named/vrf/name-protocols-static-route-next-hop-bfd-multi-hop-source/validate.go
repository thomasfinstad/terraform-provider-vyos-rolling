// Package namedvrfnameprotocolsstaticroutenexthopbfdmultihopsource code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsstaticroutenexthopbfdmultihopsource

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsStaticRouteNextHopBfdMultiHopSource{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsStaticRouteNextHopBfdMultiHopSource{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsStaticRouteNextHopBfdMultiHopSource{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsStaticRouteNextHopBfdMultiHopSource{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsStaticRouteNextHopBfdMultiHopSource{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsStaticRouteNextHopBfdMultiHopSource{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsStaticRouteNextHopBfdMultiHopSource{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsStaticRouteNextHopBfdMultiHopSource{}
