// Package namedprotocolsnhrptunneldynamicmap code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsnhrptunneldynamicmap

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsNhrpTunnelDynamicMap{}
	_ resource.ResourceWithConfigure = &protocolsNhrpTunnelDynamicMap{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsNhrpTunnelDynamicMap{}
// var _ resource.ResourceWithModifyPlan = &protocolsNhrpTunnelDynamicMap{}
// var _ resource.ResourceWithUpgradeState = &protocolsNhrpTunnelDynamicMap{}
// var _ resource.ResourceWithValidateConfig = &protocolsNhrpTunnelDynamicMap{}
// var _ resource.ResourceWithImportState = &protocolsNhrpTunnelDynamicMap{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsNhrpTunnelDynamicMap{}
