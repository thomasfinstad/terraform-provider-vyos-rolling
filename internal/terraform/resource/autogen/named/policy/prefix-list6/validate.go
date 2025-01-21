/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedpolicyprefixlistsix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyprefixlistsix

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (prefix-list6) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &policyPrefixListsix{}
	_ resource.ResourceWithConfigure   = &policyPrefixListsix{}
	_ resource.ResourceWithImportState = &policyPrefixListsix{}
)

// var _ resource.ResourceWithConfigValidators = &policyPrefixListsix{}
// var _ resource.ResourceWithModifyPlan = &policyPrefixListsix{}
// var _ resource.ResourceWithUpgradeState = &policyPrefixListsix{}
// var _ resource.ResourceWithValidateConfig = &policyPrefixListsix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyPrefixListsix{}
