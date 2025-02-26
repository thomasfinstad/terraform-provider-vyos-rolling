/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalvpnpptpremoteaccesssnmp code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalvpnpptpremoteaccesssnmp

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (snmp) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnPptpRemoteAccessSnmp{}
	_ resource.ResourceWithConfigure   = &vpnPptpRemoteAccessSnmp{}
	_ resource.ResourceWithImportState = &vpnPptpRemoteAccessSnmp{}
)

// var _ resource.ResourceWithConfigValidators = &vpnPptpRemoteAccessSnmp{}
// var _ resource.ResourceWithModifyPlan = &vpnPptpRemoteAccessSnmp{}
// var _ resource.ResourceWithUpgradeState = &vpnPptpRemoteAccessSnmp{}
// var _ resource.ResourceWithValidateConfig = &vpnPptpRemoteAccessSnmp{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnPptpRemoteAccessSnmp{}
