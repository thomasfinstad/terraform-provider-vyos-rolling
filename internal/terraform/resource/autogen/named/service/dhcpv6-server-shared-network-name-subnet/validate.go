// Package namedservicedhcpvsixserversharednetworknamesubnet code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpvsixserversharednetworknamesubnet

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceDhcpvsixServerSharedNetworkNameSubnet{}
	_ resource.ResourceWithConfigure   = &serviceDhcpvsixServerSharedNetworkNameSubnet{}
	_ resource.ResourceWithImportState = &serviceDhcpvsixServerSharedNetworkNameSubnet{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpvsixServerSharedNetworkNameSubnet{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpvsixServerSharedNetworkNameSubnet{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpvsixServerSharedNetworkNameSubnet{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpvsixServerSharedNetworkNameSubnet{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpvsixServerSharedNetworkNameSubnet{}
