/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvrfname code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfname

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfName{}
	_ resource.ResourceWithConfigure   = &vrfName{}
	_ resource.ResourceWithImportState = &vrfName{}
)

// var _ resource.ResourceWithConfigValidators = &vrfName{}
// var _ resource.ResourceWithModifyPlan = &vrfName{}
// var _ resource.ResourceWithUpgradeState = &vrfName{}
// var _ resource.ResourceWithValidateConfig = &vrfName{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfName{}
