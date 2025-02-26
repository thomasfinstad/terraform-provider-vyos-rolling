/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedservicedhcpserversharednetworkname code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicedhcpserversharednetworkname

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (shared-network-name) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceDhcpServerSharedNetworkName{}
	_ resource.ResourceWithConfigure   = &serviceDhcpServerSharedNetworkName{}
	_ resource.ResourceWithImportState = &serviceDhcpServerSharedNetworkName{}
)

// var _ resource.ResourceWithConfigValidators = &serviceDhcpServerSharedNetworkName{}
// var _ resource.ResourceWithModifyPlan = &serviceDhcpServerSharedNetworkName{}
// var _ resource.ResourceWithUpgradeState = &serviceDhcpServerSharedNetworkName{}
// var _ resource.ResourceWithValidateConfig = &serviceDhcpServerSharedNetworkName{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceDhcpServerSharedNetworkName{}
