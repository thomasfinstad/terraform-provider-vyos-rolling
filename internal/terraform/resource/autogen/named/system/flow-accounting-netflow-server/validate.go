// Package namedsystemflowaccountingnetflowserver code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemflowaccountingnetflowserver

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemFlowAccountingNetflowServer{}
	_ resource.ResourceWithConfigure = &systemFlowAccountingNetflowServer{}
)

// var _ resource.ResourceWithConfigValidators = &systemFlowAccountingNetflowServer{}
// var _ resource.ResourceWithModifyPlan = &systemFlowAccountingNetflowServer{}
// var _ resource.ResourceWithUpgradeState = &systemFlowAccountingNetflowServer{}
// var _ resource.ResourceWithValidateConfig = &systemFlowAccountingNetflowServer{}
// var _ resource.ResourceWithImportState = &systemFlowAccountingNetflowServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemFlowAccountingNetflowServer{}
