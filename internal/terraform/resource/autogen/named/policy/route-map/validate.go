/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedpolicyroutemap code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyroutemap

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (route-map) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &policyRouteMap{}
	_ resource.ResourceWithConfigure   = &policyRouteMap{}
	_ resource.ResourceWithImportState = &policyRouteMap{}
)

// var _ resource.ResourceWithConfigValidators = &policyRouteMap{}
// var _ resource.ResourceWithModifyPlan = &policyRouteMap{}
// var _ resource.ResourceWithUpgradeState = &policyRouteMap{}
// var _ resource.ResourceWithValidateConfig = &policyRouteMap{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyRouteMap{}
