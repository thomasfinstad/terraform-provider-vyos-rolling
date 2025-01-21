/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsbgpaddressfamilyltwovpnevpnroutetarget code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyltwovpnevpnroutetarget

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (route-target) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget{}
	_ resource.ResourceWithConfigure   = &protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget{}
	_ resource.ResourceWithImportState = &protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget{}
