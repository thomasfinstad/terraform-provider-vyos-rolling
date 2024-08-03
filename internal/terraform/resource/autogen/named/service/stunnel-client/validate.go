// Package namedservicestunnelclient code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicestunnelclient

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceStunnelClient{}
	_ resource.ResourceWithConfigure = &serviceStunnelClient{}
)

// var _ resource.ResourceWithConfigValidators = &serviceStunnelClient{}
// var _ resource.ResourceWithModifyPlan = &serviceStunnelClient{}
// var _ resource.ResourceWithUpgradeState = &serviceStunnelClient{}
// var _ resource.ResourceWithValidateConfig = &serviceStunnelClient{}
// var _ resource.ResourceWithImportState = &serviceStunnelClient{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceStunnelClient{}
