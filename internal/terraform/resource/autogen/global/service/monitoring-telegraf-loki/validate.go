/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalservicemonitoringtelegrafloki code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringtelegrafloki

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceMonitoringTelegrafLoki{}
	_ resource.ResourceWithConfigure   = &serviceMonitoringTelegrafLoki{}
	_ resource.ResourceWithImportState = &serviceMonitoringTelegrafLoki{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringTelegrafLoki{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringTelegrafLoki{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringTelegrafLoki{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringTelegrafLoki{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringTelegrafLoki{}
