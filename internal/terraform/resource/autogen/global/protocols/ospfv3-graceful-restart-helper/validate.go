// Package globalprotocolsospfvthreegracefulrestarthelper code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfvthreegracefulrestarthelper

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfvthreeGracefulRestartHelper{}
	_ resource.ResourceWithConfigure   = &protocolsOspfvthreeGracefulRestartHelper{}
	_ resource.ResourceWithImportState = &protocolsOspfvthreeGracefulRestartHelper{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfvthreeGracefulRestartHelper{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfvthreeGracefulRestartHelper{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfvthreeGracefulRestartHelper{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfvthreeGracefulRestartHelper{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfvthreeGracefulRestartHelper{}
