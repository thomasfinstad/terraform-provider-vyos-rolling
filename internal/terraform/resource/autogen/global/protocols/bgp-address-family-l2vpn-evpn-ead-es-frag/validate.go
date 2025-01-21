/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsbgpaddressfamilyltwovpnevpneadesfrag code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyltwovpnevpneadesfrag

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (ead-es-frag) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag{}
	_ resource.ResourceWithConfigure   = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag{}
	_ resource.ResourceWithImportState = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyLtwovpnEvpnEadEsFrag{}
