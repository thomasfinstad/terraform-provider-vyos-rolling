// Package namedservicestunnelclientpsk code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicestunnelclientpsk

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceStunnelClientPsk{}
	_ resource.ResourceWithConfigure = &serviceStunnelClientPsk{}
)

// var _ resource.ResourceWithConfigValidators = &serviceStunnelClientPsk{}
// var _ resource.ResourceWithModifyPlan = &serviceStunnelClientPsk{}
// var _ resource.ResourceWithUpgradeState = &serviceStunnelClientPsk{}
// var _ resource.ResourceWithValidateConfig = &serviceStunnelClientPsk{}
// var _ resource.ResourceWithImportState = &serviceStunnelClientPsk{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceStunnelClientPsk{}