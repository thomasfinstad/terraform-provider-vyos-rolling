// Package globalservicemonitoringtelegraf code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringtelegraf

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceMonitoringTelegraf{}
	_ resource.ResourceWithConfigure = &serviceMonitoringTelegraf{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringTelegraf{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringTelegraf{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringTelegraf{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringTelegraf{}
// var _ resource.ResourceWithImportState = &serviceMonitoringTelegraf{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringTelegraf{}