/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedpolicylocalroutesixrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicylocalroutesixrule

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &policyLocalRoutesixRule{}
	_ resource.ResourceWithConfigure   = &policyLocalRoutesixRule{}
	_ resource.ResourceWithImportState = &policyLocalRoutesixRule{}
)

// var _ resource.ResourceWithConfigValidators = &policyLocalRoutesixRule{}
// var _ resource.ResourceWithModifyPlan = &policyLocalRoutesixRule{}
// var _ resource.ResourceWithUpgradeState = &policyLocalRoutesixRule{}
// var _ resource.ResourceWithValidateConfig = &policyLocalRoutesixRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyLocalRoutesixRule{}
