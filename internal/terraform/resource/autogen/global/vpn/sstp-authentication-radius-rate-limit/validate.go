/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalvpnsstpauthenticationradiusratelimit code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnsstpauthenticationradiusratelimit

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnSstpAuthenticationRadiusRateLimit{}
	_ resource.ResourceWithConfigure   = &vpnSstpAuthenticationRadiusRateLimit{}
	_ resource.ResourceWithImportState = &vpnSstpAuthenticationRadiusRateLimit{}
)

// var _ resource.ResourceWithConfigValidators = &vpnSstpAuthenticationRadiusRateLimit{}
// var _ resource.ResourceWithModifyPlan = &vpnSstpAuthenticationRadiusRateLimit{}
// var _ resource.ResourceWithUpgradeState = &vpnSstpAuthenticationRadiusRateLimit{}
// var _ resource.ResourceWithValidateConfig = &vpnSstpAuthenticationRadiusRateLimit{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnSstpAuthenticationRadiusRateLimit{}
