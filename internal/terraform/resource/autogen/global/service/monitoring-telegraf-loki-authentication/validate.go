/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalservicemonitoringtelegraflokiauthentication code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringtelegraflokiauthentication

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceMonitoringTelegrafLokiAuthentication{}
	_ resource.ResourceWithConfigure   = &serviceMonitoringTelegrafLokiAuthentication{}
	_ resource.ResourceWithImportState = &serviceMonitoringTelegrafLokiAuthentication{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringTelegrafLokiAuthentication{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringTelegrafLokiAuthentication{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringTelegrafLokiAuthentication{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringTelegrafLokiAuthentication{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringTelegrafLokiAuthentication{}
