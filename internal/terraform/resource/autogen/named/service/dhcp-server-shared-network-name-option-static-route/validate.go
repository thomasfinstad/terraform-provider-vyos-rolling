/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedservicedhcpserversharednetworknameoptionstaticroute code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpserversharednetworknameoptionstaticroute

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}
	_ resource.ResourceWithConfigure   = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}
	_ resource.ResourceWithImportState = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpServerSharedNetworkNameOptionStaticRoute{}
