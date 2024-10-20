/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalvpnsstp code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnsstp

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnSstp{}
	_ resource.ResourceWithConfigure   = &vpnSstp{}
	_ resource.ResourceWithImportState = &vpnSstp{}
)

// var _ resource.ResourceWithConfigValidators = &vpnSstp{}
// var _ resource.ResourceWithModifyPlan = &vpnSstp{}
// var _ resource.ResourceWithUpgradeState = &vpnSstp{}
// var _ resource.ResourceWithValidateConfig = &vpnSstp{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnSstp{}
