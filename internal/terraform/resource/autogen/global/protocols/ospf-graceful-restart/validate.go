// Package globalprotocolsospfgracefulrestart code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfgracefulrestart

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfGracefulRestart{}
	_ resource.ResourceWithConfigure = &protocolsOspfGracefulRestart{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfGracefulRestart{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfGracefulRestart{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfGracefulRestart{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfGracefulRestart{}
// var _ resource.ResourceWithImportState = &protocolsOspfGracefulRestart{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfGracefulRestart{}
