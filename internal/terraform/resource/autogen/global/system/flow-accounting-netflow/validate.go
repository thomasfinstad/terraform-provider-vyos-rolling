/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalsystemflowaccountingnetflow code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemflowaccountingnetflow

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (netflow) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemFlowAccountingNetflow{}
	_ resource.ResourceWithConfigure   = &systemFlowAccountingNetflow{}
	_ resource.ResourceWithImportState = &systemFlowAccountingNetflow{}
)

// var _ resource.ResourceWithConfigValidators = &systemFlowAccountingNetflow{}
// var _ resource.ResourceWithModifyPlan = &systemFlowAccountingNetflow{}
// var _ resource.ResourceWithUpgradeState = &systemFlowAccountingNetflow{}
// var _ resource.ResourceWithValidateConfig = &systemFlowAccountingNetflow{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemFlowAccountingNetflow{}
