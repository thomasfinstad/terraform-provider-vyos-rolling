/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvpnsstpauthenticationlocalusersusername code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnsstpauthenticationlocalusersusername

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnSstpAuthenticationLocalUsersUsername{}
	_ resource.ResourceWithConfigure   = &vpnSstpAuthenticationLocalUsersUsername{}
	_ resource.ResourceWithImportState = &vpnSstpAuthenticationLocalUsersUsername{}
)

// var _ resource.ResourceWithConfigValidators = &vpnSstpAuthenticationLocalUsersUsername{}
// var _ resource.ResourceWithModifyPlan = &vpnSstpAuthenticationLocalUsersUsername{}
// var _ resource.ResourceWithUpgradeState = &vpnSstpAuthenticationLocalUsersUsername{}
// var _ resource.ResourceWithValidateConfig = &vpnSstpAuthenticationLocalUsersUsername{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnSstpAuthenticationLocalUsersUsername{}
