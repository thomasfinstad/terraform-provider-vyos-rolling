// Package globalservicemonitoringzabbixagentlimits code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringzabbixagentlimits

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceMonitoringZabbixAgentLimits{}
	_ resource.ResourceWithConfigure = &serviceMonitoringZabbixAgentLimits{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringZabbixAgentLimits{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringZabbixAgentLimits{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringZabbixAgentLimits{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringZabbixAgentLimits{}
// var _ resource.ResourceWithImportState = &serviceMonitoringZabbixAgentLimits{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringZabbixAgentLimits{}