// Package namedservicedhcpserversharednetworknamesubnetstaticmapping code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpserversharednetworknamesubnetstaticmapping

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceDhcpServerSharedNetworkNameSubnetStaticMapping{}
	_ resource.ResourceWithConfigure   = &serviceDhcpServerSharedNetworkNameSubnetStaticMapping{}
	_ resource.ResourceWithImportState = &serviceDhcpServerSharedNetworkNameSubnetStaticMapping{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpServerSharedNetworkNameSubnetStaticMapping{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpServerSharedNetworkNameSubnetStaticMapping{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpServerSharedNetworkNameSubnetStaticMapping{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpServerSharedNetworkNameSubnetStaticMapping{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpServerSharedNetworkNameSubnetStaticMapping{}
