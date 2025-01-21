/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvpnltwotpremoteaccessclientipvsixpool code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnltwotpremoteaccessclientipvsixpool

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (client-ipv6-pool) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnLtwotpRemoteAccessClientIPvsixPool{}
	_ resource.ResourceWithConfigure   = &vpnLtwotpRemoteAccessClientIPvsixPool{}
	_ resource.ResourceWithImportState = &vpnLtwotpRemoteAccessClientIPvsixPool{}
)

// var _ resource.ResourceWithConfigValidators = &vpnLtwotpRemoteAccessClientIPvsixPool{}
// var _ resource.ResourceWithModifyPlan = &vpnLtwotpRemoteAccessClientIPvsixPool{}
// var _ resource.ResourceWithUpgradeState = &vpnLtwotpRemoteAccessClientIPvsixPool{}
// var _ resource.ResourceWithValidateConfig = &vpnLtwotpRemoteAccessClientIPvsixPool{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnLtwotpRemoteAccessClientIPvsixPool{}
