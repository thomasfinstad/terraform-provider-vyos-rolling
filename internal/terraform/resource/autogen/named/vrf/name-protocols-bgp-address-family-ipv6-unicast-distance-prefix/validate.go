/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvrfnameprotocolsbgpaddressfamilyipvsixunicastdistanceprefix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvsixunicastdistanceprefix

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (prefix) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpAddressFamilyIPvsixUnicastDistancePrefix{}
