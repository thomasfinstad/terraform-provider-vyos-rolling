// Package namedservicedhcpvsixserversharednetworknamesubnetprefixdelegationprefix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpvsixserversharednetworknamesubnetprefixdelegationprefix

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceDhcpvsixServerSharedNetworkNameSubnetPrefixDelegationPrefix{}
	_ resource.ResourceWithConfigure   = &serviceDhcpvsixServerSharedNetworkNameSubnetPrefixDelegationPrefix{}
	_ resource.ResourceWithImportState = &serviceDhcpvsixServerSharedNetworkNameSubnetPrefixDelegationPrefix{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpvsixServerSharedNetworkNameSubnetPrefixDelegationPrefix{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpvsixServerSharedNetworkNameSubnetPrefixDelegationPrefix{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpvsixServerSharedNetworkNameSubnetPrefixDelegationPrefix{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpvsixServerSharedNetworkNameSubnetPrefixDelegationPrefix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpvsixServerSharedNetworkNameSubnetPrefixDelegationPrefix{}
