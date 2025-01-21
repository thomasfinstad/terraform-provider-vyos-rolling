/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedpolicyprefixlistsixrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyprefixlistsixrule

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (rule) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &policyPrefixListsixRule{}
	_ resource.ResourceWithConfigure   = &policyPrefixListsixRule{}
	_ resource.ResourceWithImportState = &policyPrefixListsixRule{}
)

// var _ resource.ResourceWithConfigValidators = &policyPrefixListsixRule{}
// var _ resource.ResourceWithModifyPlan = &policyPrefixListsixRule{}
// var _ resource.ResourceWithUpgradeState = &policyPrefixListsixRule{}
// var _ resource.ResourceWithValidateConfig = &policyPrefixListsixRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyPrefixListsixRule{}
