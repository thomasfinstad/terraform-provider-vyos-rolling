// Package namedservicedhcpserversharednetworknameoptionstaticroute code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpserversharednetworknameoptionstaticroute

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}
	_ resource.ResourceWithConfigure = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}
// var _ resource.ResourceWithImportState = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}
