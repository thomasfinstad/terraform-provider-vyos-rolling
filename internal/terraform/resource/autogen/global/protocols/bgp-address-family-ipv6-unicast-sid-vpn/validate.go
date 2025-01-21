/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsbgpaddressfamilyipvsixunicastsidvpn code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvsixunicastsidvpn

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (vpn) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpAddressFamilyIPvsixUnicastSIDVpn{}
	_ resource.ResourceWithConfigure   = &protocolsBgpAddressFamilyIPvsixUnicastSIDVpn{}
	_ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvsixUnicastSIDVpn{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvsixUnicastSIDVpn{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvsixUnicastSIDVpn{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvsixUnicastSIDVpn{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvsixUnicastSIDVpn{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvsixUnicastSIDVpn{}
