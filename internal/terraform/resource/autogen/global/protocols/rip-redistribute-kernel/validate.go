/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsripredistributekernel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsripredistributekernel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (kernel) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsRIPRedistributeKernel{}
	_ resource.ResourceWithConfigure   = &protocolsRIPRedistributeKernel{}
	_ resource.ResourceWithImportState = &protocolsRIPRedistributeKernel{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPRedistributeKernel{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPRedistributeKernel{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPRedistributeKernel{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPRedistributeKernel{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPRedistributeKernel{}
