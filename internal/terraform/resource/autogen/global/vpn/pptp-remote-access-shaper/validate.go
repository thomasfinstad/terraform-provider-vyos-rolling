/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalvpnpptpremoteaccessshaper code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnpptpremoteaccessshaper

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (shaper) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnPptpRemoteAccessShaper{}
	_ resource.ResourceWithConfigure   = &vpnPptpRemoteAccessShaper{}
	_ resource.ResourceWithImportState = &vpnPptpRemoteAccessShaper{}
)

// var _ resource.ResourceWithConfigValidators = &vpnPptpRemoteAccessShaper{}
// var _ resource.ResourceWithModifyPlan = &vpnPptpRemoteAccessShaper{}
// var _ resource.ResourceWithUpgradeState = &vpnPptpRemoteAccessShaper{}
// var _ resource.ResourceWithValidateConfig = &vpnPptpRemoteAccessShaper{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnPptpRemoteAccessShaper{}
