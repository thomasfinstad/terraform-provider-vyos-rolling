// Package globalsystemflowaccountingnetflowtimeout code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemflowaccountingnetflowtimeout

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemFlowAccountingNetflowTimeout{}
	_ resource.ResourceWithConfigure = &systemFlowAccountingNetflowTimeout{}
)

// var _ resource.ResourceWithConfigValidators = &systemFlowAccountingNetflowTimeout{}
// var _ resource.ResourceWithModifyPlan = &systemFlowAccountingNetflowTimeout{}
// var _ resource.ResourceWithUpgradeState = &systemFlowAccountingNetflowTimeout{}
// var _ resource.ResourceWithValidateConfig = &systemFlowAccountingNetflowTimeout{}
// var _ resource.ResourceWithImportState = &systemFlowAccountingNetflowTimeout{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemFlowAccountingNetflowTimeout{}
