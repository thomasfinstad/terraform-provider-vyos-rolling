/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedpolicyroute code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyroute

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &policyRoute{}
	_ resource.ResourceWithConfigure   = &policyRoute{}
	_ resource.ResourceWithImportState = &policyRoute{}
)

// var _ resource.ResourceWithConfigValidators = &policyRoute{}
// var _ resource.ResourceWithModifyPlan = &policyRoute{}
// var _ resource.ResourceWithUpgradeState = &policyRoute{}
// var _ resource.ResourceWithValidateConfig = &policyRoute{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyRoute{}
