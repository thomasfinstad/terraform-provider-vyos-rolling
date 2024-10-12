// Package globalprotocolsospfmaxmetricrouterlsa code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsospfmaxmetricrouterlsa

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsOspfMaxMetricRouterLsa{}
	_ resource.ResourceWithConfigure = &protocolsOspfMaxMetricRouterLsa{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsOspfMaxMetricRouterLsa{}
// var _ resource.ResourceWithModifyPlan = &protocolsOspfMaxMetricRouterLsa{}
// var _ resource.ResourceWithUpgradeState = &protocolsOspfMaxMetricRouterLsa{}
// var _ resource.ResourceWithValidateConfig = &protocolsOspfMaxMetricRouterLsa{}
// var _ resource.ResourceWithImportState = &protocolsOspfMaxMetricRouterLsa{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsOspfMaxMetricRouterLsa{}
