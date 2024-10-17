// Package globalservicemonitoringtelegrafinfluxdbauthentication code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringtelegrafinfluxdbauthentication

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceMonitoringTelegrafInfluxdbAuthentication{}
	_ resource.ResourceWithConfigure   = &serviceMonitoringTelegrafInfluxdbAuthentication{}
	_ resource.ResourceWithImportState = &serviceMonitoringTelegrafInfluxdbAuthentication{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringTelegrafInfluxdbAuthentication{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringTelegrafInfluxdbAuthentication{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringTelegrafInfluxdbAuthentication{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringTelegrafInfluxdbAuthentication{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringTelegrafInfluxdbAuthentication{}
