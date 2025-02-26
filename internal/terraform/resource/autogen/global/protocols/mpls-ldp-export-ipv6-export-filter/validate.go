/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsmplsldpexportipvsixexportfilter code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsmplsldpexportipvsixexportfilter

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (export-filter) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsMplsLdpExportIPvsixExportFilter{}
	_ resource.ResourceWithConfigure   = &protocolsMplsLdpExportIPvsixExportFilter{}
	_ resource.ResourceWithImportState = &protocolsMplsLdpExportIPvsixExportFilter{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsMplsLdpExportIPvsixExportFilter{}
// var _ resource.ResourceWithModifyPlan = &protocolsMplsLdpExportIPvsixExportFilter{}
// var _ resource.ResourceWithUpgradeState = &protocolsMplsLdpExportIPvsixExportFilter{}
// var _ resource.ResourceWithValidateConfig = &protocolsMplsLdpExportIPvsixExportFilter{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsMplsLdpExportIPvsixExportFilter{}
