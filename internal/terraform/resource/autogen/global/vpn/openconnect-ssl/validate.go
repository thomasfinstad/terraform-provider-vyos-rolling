/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalvpnopenconnectssl code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnopenconnectssl

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (ssl) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnOpenconnectSsl{}
	_ resource.ResourceWithConfigure   = &vpnOpenconnectSsl{}
	_ resource.ResourceWithImportState = &vpnOpenconnectSsl{}
)

// var _ resource.ResourceWithConfigValidators = &vpnOpenconnectSsl{}
// var _ resource.ResourceWithModifyPlan = &vpnOpenconnectSsl{}
// var _ resource.ResourceWithUpgradeState = &vpnOpenconnectSsl{}
// var _ resource.ResourceWithValidateConfig = &vpnOpenconnectSsl{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnOpenconnectSsl{}
