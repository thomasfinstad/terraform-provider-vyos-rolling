/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedpolicyprefixlist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyprefixlist

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (prefix-list) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &policyPrefixList{}
	_ resource.ResourceWithConfigure   = &policyPrefixList{}
	_ resource.ResourceWithImportState = &policyPrefixList{}
)

// var _ resource.ResourceWithConfigValidators = &policyPrefixList{}
// var _ resource.ResourceWithModifyPlan = &policyPrefixList{}
// var _ resource.ResourceWithUpgradeState = &policyPrefixList{}
// var _ resource.ResourceWithValidateConfig = &policyPrefixList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyPrefixList{}
