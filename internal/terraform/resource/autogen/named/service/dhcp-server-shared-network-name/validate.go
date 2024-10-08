// Package namedservicedhcpserversharednetworkname code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpserversharednetworkname

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceDhcpServerSharedNetworkName{}
	_ resource.ResourceWithConfigure = &serviceDhcpServerSharedNetworkName{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpServerSharedNetworkName{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpServerSharedNetworkName{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpServerSharedNetworkName{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpServerSharedNetworkName{}
// var _ resource.ResourceWithImportState = &serviceDhcpServerSharedNetworkName{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpServerSharedNetworkName{}
