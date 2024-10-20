/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalservicemonitoringzabbixagentlimits code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringzabbixagentlimits

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceMonitoringZabbixAgentLimits{}
	_ resource.ResourceWithConfigure   = &serviceMonitoringZabbixAgentLimits{}
	_ resource.ResourceWithImportState = &serviceMonitoringZabbixAgentLimits{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringZabbixAgentLimits{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringZabbixAgentLimits{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringZabbixAgentLimits{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringZabbixAgentLimits{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringZabbixAgentLimits{}
