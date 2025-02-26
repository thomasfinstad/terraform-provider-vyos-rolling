/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvpnsstpclientipvsixpool code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnsstpclientipvsixpool

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (client-ipv6-pool) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnSstpClientIPvsixPool{}
	_ resource.ResourceWithConfigure   = &vpnSstpClientIPvsixPool{}
	_ resource.ResourceWithImportState = &vpnSstpClientIPvsixPool{}
)

// var _ resource.ResourceWithConfigValidators = &vpnSstpClientIPvsixPool{}
// var _ resource.ResourceWithModifyPlan = &vpnSstpClientIPvsixPool{}
// var _ resource.ResourceWithUpgradeState = &vpnSstpClientIPvsixPool{}
// var _ resource.ResourceWithValidateConfig = &vpnSstpClientIPvsixPool{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnSstpClientIPvsixPool{}
