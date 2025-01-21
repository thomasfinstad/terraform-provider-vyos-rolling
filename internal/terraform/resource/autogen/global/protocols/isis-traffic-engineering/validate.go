/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsisistrafficengineering code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisistrafficengineering

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (traffic-engineering) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsIsisTrafficEngineering{}
	_ resource.ResourceWithConfigure   = &protocolsIsisTrafficEngineering{}
	_ resource.ResourceWithImportState = &protocolsIsisTrafficEngineering{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisTrafficEngineering{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisTrafficEngineering{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisTrafficEngineering{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisTrafficEngineering{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisTrafficEngineering{}
