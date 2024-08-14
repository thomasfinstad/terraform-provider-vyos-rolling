// Package globalservicemonitoringtelegrafazuredataexplorerauthentication code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringtelegrafazuredataexplorerauthentication

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}
	_ resource.ResourceWithConfigure = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}
// var _ resource.ResourceWithImportState = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}