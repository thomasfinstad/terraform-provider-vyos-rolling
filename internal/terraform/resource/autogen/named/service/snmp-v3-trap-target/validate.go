/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedservicesnmpvthreetraptarget code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicesnmpvthreetraptarget

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (trap-target) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceSnmpVthreeTrapTarget{}
	_ resource.ResourceWithConfigure   = &serviceSnmpVthreeTrapTarget{}
	_ resource.ResourceWithImportState = &serviceSnmpVthreeTrapTarget{}
)

// var _ resource.ResourceWithConfigValidators = &serviceSnmpVthreeTrapTarget{}
// var _ resource.ResourceWithModifyPlan = &serviceSnmpVthreeTrapTarget{}
// var _ resource.ResourceWithUpgradeState = &serviceSnmpVthreeTrapTarget{}
// var _ resource.ResourceWithValidateConfig = &serviceSnmpVthreeTrapTarget{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceSnmpVthreeTrapTarget{}
