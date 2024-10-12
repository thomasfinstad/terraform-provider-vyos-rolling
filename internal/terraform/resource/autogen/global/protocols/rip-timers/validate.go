// Package globalprotocolsriptimers code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsriptimers

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsRIPTimers{}
	_ resource.ResourceWithConfigure = &protocolsRIPTimers{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsRIPTimers{}
// var _ resource.ResourceWithModifyPlan = &protocolsRIPTimers{}
// var _ resource.ResourceWithUpgradeState = &protocolsRIPTimers{}
// var _ resource.ResourceWithValidateConfig = &protocolsRIPTimers{}
// var _ resource.ResourceWithImportState = &protocolsRIPTimers{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsRIPTimers{}
