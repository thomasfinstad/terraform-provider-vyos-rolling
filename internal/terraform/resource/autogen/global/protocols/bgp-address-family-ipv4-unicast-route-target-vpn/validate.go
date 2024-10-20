/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsbgpaddressfamilyipvfourunicastroutetargetvpn code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvfourunicastroutetargetvpn

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}
	_ resource.ResourceWithConfigure   = &protocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}
	_ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn{}
