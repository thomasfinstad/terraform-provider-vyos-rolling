/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsbgpaddressfamilyltwovpnevpneadesroutetarget code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyltwovpnevpneadesroutetarget

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (ead-es-route-target) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget{}
	_ resource.ResourceWithConfigure   = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget{}
	_ resource.ResourceWithImportState = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsRouteTarget{}
