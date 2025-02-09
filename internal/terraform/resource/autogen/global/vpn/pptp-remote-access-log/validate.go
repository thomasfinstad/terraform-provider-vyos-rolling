/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalvpnpptpremoteaccesslog code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnpptpremoteaccesslog

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (log) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnPptpRemoteAccessLog{}
	_ resource.ResourceWithConfigure   = &vpnPptpRemoteAccessLog{}
	_ resource.ResourceWithImportState = &vpnPptpRemoteAccessLog{}
)

// var _ resource.ResourceWithConfigValidators = &vpnPptpRemoteAccessLog{}
// var _ resource.ResourceWithModifyPlan = &vpnPptpRemoteAccessLog{}
// var _ resource.ResourceWithUpgradeState = &vpnPptpRemoteAccessLog{}
// var _ resource.ResourceWithValidateConfig = &vpnPptpRemoteAccessLog{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnPptpRemoteAccessLog{}
