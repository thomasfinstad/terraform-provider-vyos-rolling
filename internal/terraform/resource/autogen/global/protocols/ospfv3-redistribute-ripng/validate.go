/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsospfvthreeredistributeripng code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfvthreeredistributeripng

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (ripng) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfvthreeRedistributeRIPng{}
	_ resource.ResourceWithConfigure   = &protocolsOspfvthreeRedistributeRIPng{}
	_ resource.ResourceWithImportState = &protocolsOspfvthreeRedistributeRIPng{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfvthreeRedistributeRIPng{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfvthreeRedistributeRIPng{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfvthreeRedistributeRIPng{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfvthreeRedistributeRIPng{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfvthreeRedistributeRIPng{}
