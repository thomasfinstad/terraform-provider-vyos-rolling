// Package namedservicestunnelserverpsk code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicestunnelserverpsk

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceStunnelServerPsk{}
	_ resource.ResourceWithConfigure = &serviceStunnelServerPsk{}
)

// var _ resource.ResourceWithConfigValidators = &serviceStunnelServerPsk{}
// var _ resource.ResourceWithModifyPlan = &serviceStunnelServerPsk{}
// var _ resource.ResourceWithUpgradeState = &serviceStunnelServerPsk{}
// var _ resource.ResourceWithValidateConfig = &serviceStunnelServerPsk{}
// var _ resource.ResourceWithImportState = &serviceStunnelServerPsk{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceStunnelServerPsk{}
