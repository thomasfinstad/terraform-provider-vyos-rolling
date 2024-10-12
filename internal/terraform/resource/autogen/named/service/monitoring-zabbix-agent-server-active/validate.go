// Package namedservicemonitoringzabbixagentserveractive code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicemonitoringzabbixagentserveractive

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceMonitoringZabbixAgentServerActive{}
	_ resource.ResourceWithConfigure = &serviceMonitoringZabbixAgentServerActive{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringZabbixAgentServerActive{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringZabbixAgentServerActive{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringZabbixAgentServerActive{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringZabbixAgentServerActive{}
// var _ resource.ResourceWithImportState = &serviceMonitoringZabbixAgentServerActive{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringZabbixAgentServerActive{}
