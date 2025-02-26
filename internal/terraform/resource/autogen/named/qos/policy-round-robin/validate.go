/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedqospolicyroundrobin code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicyroundrobin

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (round-robin) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &qosPolicyRoundRobin{}
	_ resource.ResourceWithConfigure   = &qosPolicyRoundRobin{}
	_ resource.ResourceWithImportState = &qosPolicyRoundRobin{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyRoundRobin{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyRoundRobin{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyRoundRobin{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyRoundRobin{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyRoundRobin{}
