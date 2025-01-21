/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicemonitoringtelegrafazuredataexplorerauthentication code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringtelegrafazuredataexplorerauthentication

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (authentication) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}
	_ resource.ResourceWithConfigure   = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}
	_ resource.ResourceWithImportState = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringTelegrafAzureDataExplorerAuthentication{}
