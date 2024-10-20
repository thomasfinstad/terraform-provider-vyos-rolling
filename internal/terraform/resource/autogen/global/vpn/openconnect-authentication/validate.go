/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalvpnopenconnectauthentication code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnopenconnectauthentication

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnOpenconnectAuthentication{}
	_ resource.ResourceWithConfigure   = &vpnOpenconnectAuthentication{}
	_ resource.ResourceWithImportState = &vpnOpenconnectAuthentication{}
)

// var _ resource.ResourceWithConfigValidators = &vpnOpenconnectAuthentication{}
// var _ resource.ResourceWithModifyPlan = &vpnOpenconnectAuthentication{}
// var _ resource.ResourceWithUpgradeState = &vpnOpenconnectAuthentication{}
// var _ resource.ResourceWithValidateConfig = &vpnOpenconnectAuthentication{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnOpenconnectAuthentication{}
