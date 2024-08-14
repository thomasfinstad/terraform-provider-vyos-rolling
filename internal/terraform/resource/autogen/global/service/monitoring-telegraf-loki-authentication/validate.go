// Package globalservicemonitoringtelegraflokiauthentication code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringtelegraflokiauthentication

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceMonitoringTelegrafLokiAuthentication{}
	_ resource.ResourceWithConfigure = &serviceMonitoringTelegrafLokiAuthentication{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringTelegrafLokiAuthentication{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringTelegrafLokiAuthentication{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringTelegrafLokiAuthentication{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringTelegrafLokiAuthentication{}
// var _ resource.ResourceWithImportState = &serviceMonitoringTelegrafLokiAuthentication{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringTelegrafLokiAuthentication{}