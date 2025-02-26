/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedservicedhcpserversharednetworknamesubnetstaticmappingoptionstaticroute code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpserversharednetworknamesubnetstaticmappingoptionstaticroute

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (static-route) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}
	_ resource.ResourceWithConfigure   = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}
	_ resource.ResourceWithImportState = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute{}
