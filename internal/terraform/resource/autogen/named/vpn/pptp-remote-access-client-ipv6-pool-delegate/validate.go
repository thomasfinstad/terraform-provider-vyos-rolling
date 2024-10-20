/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvpnpptpremoteaccessclientipvsixpooldelegate code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnpptpremoteaccessclientipvsixpooldelegate

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnPptpRemoteAccessClientIPvsixPoolDelegate{}
	_ resource.ResourceWithConfigure   = &vpnPptpRemoteAccessClientIPvsixPoolDelegate{}
	_ resource.ResourceWithImportState = &vpnPptpRemoteAccessClientIPvsixPoolDelegate{}
)

// var _ resource.ResourceWithConfigValidators = &vpnPptpRemoteAccessClientIPvsixPoolDelegate{}
// var _ resource.ResourceWithModifyPlan = &vpnPptpRemoteAccessClientIPvsixPoolDelegate{}
// var _ resource.ResourceWithUpgradeState = &vpnPptpRemoteAccessClientIPvsixPoolDelegate{}
// var _ resource.ResourceWithValidateConfig = &vpnPptpRemoteAccessClientIPvsixPoolDelegate{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnPptpRemoteAccessClientIPvsixPoolDelegate{}
