/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedcontainernamelabel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedcontainernamelabel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (label) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &containerNameLabel{}
	_ resource.ResourceWithConfigure   = &containerNameLabel{}
	_ resource.ResourceWithImportState = &containerNameLabel{}
)

// var _ resource.ResourceWithConfigValidators = &containerNameLabel{}
// var _ resource.ResourceWithModifyPlan = &containerNameLabel{}
// var _ resource.ResourceWithUpgradeState = &containerNameLabel{}
// var _ resource.ResourceWithValidateConfig = &containerNameLabel{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &containerNameLabel{}
