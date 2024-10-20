/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsbgpaddressfamilyipvsixunicastredistribute code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvsixunicastredistribute

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpAddressFamilyIPvsixUnicastRedistribute{}
	_ resource.ResourceWithConfigure   = &protocolsBgpAddressFamilyIPvsixUnicastRedistribute{}
	_ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvsixUnicastRedistribute{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvsixUnicastRedistribute{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvsixUnicastRedistribute{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvsixUnicastRedistribute{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvsixUnicastRedistribute{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvsixUnicastRedistribute{}
