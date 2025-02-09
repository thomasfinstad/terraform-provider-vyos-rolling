/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsospfarea code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfarea

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (area) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfArea{}
	_ resource.ResourceWithConfigure   = &protocolsOspfArea{}
	_ resource.ResourceWithImportState = &protocolsOspfArea{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfArea{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfArea{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfArea{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfArea{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfArea{}
