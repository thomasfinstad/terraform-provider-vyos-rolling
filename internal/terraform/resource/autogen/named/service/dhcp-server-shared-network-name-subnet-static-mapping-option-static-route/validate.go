// Package namedservicedhcpserversharednetworknamesubnetstaticmappingoptionstaticroute code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpserversharednetworknamesubnetstaticmappingoptionstaticroute

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}
	_ resource.ResourceWithConfigure = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}
// var _ resource.ResourceWithImportState = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}
