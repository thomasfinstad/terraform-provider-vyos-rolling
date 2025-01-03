/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedservicemonitoringprometheusblackboxexportermodulesdnsname code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicemonitoringprometheusblackboxexportermodulesdnsname

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceMonitoringPrometheusBlackboxExporterModulesDNSName{}
	_ resource.ResourceWithConfigure   = &serviceMonitoringPrometheusBlackboxExporterModulesDNSName{}
	_ resource.ResourceWithImportState = &serviceMonitoringPrometheusBlackboxExporterModulesDNSName{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringPrometheusBlackboxExporterModulesDNSName{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringPrometheusBlackboxExporterModulesDNSName{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringPrometheusBlackboxExporterModulesDNSName{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringPrometheusBlackboxExporterModulesDNSName{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringPrometheusBlackboxExporterModulesDNSName{}
