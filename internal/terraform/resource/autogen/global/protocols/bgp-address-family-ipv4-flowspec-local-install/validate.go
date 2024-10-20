/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolsbgpaddressfamilyipvfourflowspeclocalinstall code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvfourflowspeclocalinstall

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpAddressFamilyIPvfourFlowspecLocalInstall{}
	_ resource.ResourceWithConfigure   = &protocolsBgpAddressFamilyIPvfourFlowspecLocalInstall{}
	_ resource.ResourceWithImportState = &protocolsBgpAddressFamilyIPvfourFlowspecLocalInstall{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpAddressFamilyIPvfourFlowspecLocalInstall{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpAddressFamilyIPvfourFlowspecLocalInstall{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpAddressFamilyIPvfourFlowspecLocalInstall{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpAddressFamilyIPvfourFlowspecLocalInstall{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpAddressFamilyIPvfourFlowspecLocalInstall{}
