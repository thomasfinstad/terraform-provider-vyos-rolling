// Package globalservicemonitoringnodeexporter code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringnodeexporter

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceMonitoringNodeExporter{}
	_ resource.ResourceWithConfigure = &serviceMonitoringNodeExporter{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringNodeExporter{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringNodeExporter{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringNodeExporter{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringNodeExporter{}
// var _ resource.ResourceWithImportState = &serviceMonitoringNodeExporter{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringNodeExporter{}
