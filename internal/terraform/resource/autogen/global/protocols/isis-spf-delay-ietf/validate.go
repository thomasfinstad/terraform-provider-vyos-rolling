/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsisisspfdelayietf code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisspfdelayietf

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (spf-delay-ietf) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsIsisSpfDelayIetf{}
	_ resource.ResourceWithConfigure   = &protocolsIsisSpfDelayIetf{}
	_ resource.ResourceWithImportState = &protocolsIsisSpfDelayIetf{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisSpfDelayIetf{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisSpfDelayIetf{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisSpfDelayIetf{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisSpfDelayIetf{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisSpfDelayIetf{}
