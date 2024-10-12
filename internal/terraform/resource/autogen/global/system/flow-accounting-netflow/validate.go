// Package globalsystemflowaccountingnetflow code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemflowaccountingnetflow

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemFlowAccountingNetflow{}
	_ resource.ResourceWithConfigure = &systemFlowAccountingNetflow{}
)

// var _ resource.ResourceWithConfigValidators = &systemFlowAccountingNetflow{}
// var _ resource.ResourceWithModifyPlan = &systemFlowAccountingNetflow{}
// var _ resource.ResourceWithUpgradeState = &systemFlowAccountingNetflow{}
// var _ resource.ResourceWithValidateConfig = &systemFlowAccountingNetflow{}
// var _ resource.ResourceWithImportState = &systemFlowAccountingNetflow{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemFlowAccountingNetflow{}
