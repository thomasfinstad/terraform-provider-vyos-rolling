/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicemonitoringprometheusnodeexporter code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringprometheusnodeexporter

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (node-exporter) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceMonitoringPrometheusNodeExporter{}
	_ resource.ResourceWithConfigure   = &serviceMonitoringPrometheusNodeExporter{}
	_ resource.ResourceWithImportState = &serviceMonitoringPrometheusNodeExporter{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringPrometheusNodeExporter{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringPrometheusNodeExporter{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringPrometheusNodeExporter{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringPrometheusNodeExporter{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringPrometheusNodeExporter{}
