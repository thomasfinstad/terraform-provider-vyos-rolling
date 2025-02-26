/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsospfredistributekernel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfredistributekernel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (kernel) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsOspfRedistributeKernel{}
	_ resource.ResourceWithConfigure   = &protocolsOspfRedistributeKernel{}
	_ resource.ResourceWithImportState = &protocolsOspfRedistributeKernel{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfRedistributeKernel{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfRedistributeKernel{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfRedistributeKernel{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfRedistributeKernel{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfRedistributeKernel{}
