/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalvpnpptpremoteaccessauthenticationradiusratelimit code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnpptpremoteaccessauthenticationradiusratelimit

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (rate-limit) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}
	_ resource.ResourceWithConfigure   = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}
	_ resource.ResourceWithImportState = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}
)

// var _ resource.ResourceWithConfigValidators = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}
// var _ resource.ResourceWithModifyPlan = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}
// var _ resource.ResourceWithUpgradeState = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}
// var _ resource.ResourceWithValidateConfig = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnPptpRemoteAccessAuthenticationRadiusRateLimit{}
