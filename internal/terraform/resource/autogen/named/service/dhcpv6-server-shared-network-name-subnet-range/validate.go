// Package namedservicedhcpvsixserversharednetworknamesubnetrange code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpvsixserversharednetworknamesubnetrange

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceDhcpvsixServerSharedNetworkNameSubnetRange{}
	_ resource.ResourceWithConfigure = &serviceDhcpvsixServerSharedNetworkNameSubnetRange{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpvsixServerSharedNetworkNameSubnetRange{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpvsixServerSharedNetworkNameSubnetRange{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpvsixServerSharedNetworkNameSubnetRange{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpvsixServerSharedNetworkNameSubnetRange{}
// var _ resource.ResourceWithImportState = &serviceDhcpvsixServerSharedNetworkNameSubnetRange{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpvsixServerSharedNetworkNameSubnetRange{}
