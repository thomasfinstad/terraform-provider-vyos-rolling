/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedservicesnmpvthreeview code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicesnmpvthreeview

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (view) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceSnmpVthreeView{}
	_ resource.ResourceWithConfigure   = &serviceSnmpVthreeView{}
	_ resource.ResourceWithImportState = &serviceSnmpVthreeView{}
)

// var _ resource.ResourceWithConfigValidators = &serviceSnmpVthreeView{}
// var _ resource.ResourceWithModifyPlan = &serviceSnmpVthreeView{}
// var _ resource.ResourceWithUpgradeState = &serviceSnmpVthreeView{}
// var _ resource.ResourceWithValidateConfig = &serviceSnmpVthreeView{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceSnmpVthreeView{}
