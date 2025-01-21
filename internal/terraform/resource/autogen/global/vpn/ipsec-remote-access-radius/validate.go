/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalvpnipsecremoteaccessradius code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnipsecremoteaccessradius

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (radius) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnIPsecRemoteAccessRadius{}
	_ resource.ResourceWithConfigure   = &vpnIPsecRemoteAccessRadius{}
	_ resource.ResourceWithImportState = &vpnIPsecRemoteAccessRadius{}
)

// var _ resource.ResourceWithConfigValidators = &vpnIPsecRemoteAccessRadius{}
// var _ resource.ResourceWithModifyPlan = &vpnIPsecRemoteAccessRadius{}
// var _ resource.ResourceWithUpgradeState = &vpnIPsecRemoteAccessRadius{}
// var _ resource.ResourceWithValidateConfig = &vpnIPsecRemoteAccessRadius{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnIPsecRemoteAccessRadius{}
