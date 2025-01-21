/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicemonitoringtelegrafsplunk code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringtelegrafsplunk

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (splunk) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceMonitoringTelegrafSplunk{}
	_ resource.ResourceWithConfigure   = &serviceMonitoringTelegrafSplunk{}
	_ resource.ResourceWithImportState = &serviceMonitoringTelegrafSplunk{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringTelegrafSplunk{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringTelegrafSplunk{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringTelegrafSplunk{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringTelegrafSplunk{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringTelegrafSplunk{}
