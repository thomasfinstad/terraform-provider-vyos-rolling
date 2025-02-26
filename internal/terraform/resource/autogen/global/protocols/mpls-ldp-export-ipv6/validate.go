/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsmplsldpexportipvsix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsmplsldpexportipvsix

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (ipv6) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsMplsLdpExportIPvsix{}
	_ resource.ResourceWithConfigure   = &protocolsMplsLdpExportIPvsix{}
	_ resource.ResourceWithImportState = &protocolsMplsLdpExportIPvsix{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsMplsLdpExportIPvsix{}
// var _ resource.ResourceWithModifyPlan = &protocolsMplsLdpExportIPvsix{}
// var _ resource.ResourceWithUpgradeState = &protocolsMplsLdpExportIPvsix{}
// var _ resource.ResourceWithValidateConfig = &protocolsMplsLdpExportIPvsix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsMplsLdpExportIPvsix{}
