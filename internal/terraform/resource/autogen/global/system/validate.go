/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalsystem code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystem

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &system{}
	_ resource.ResourceWithConfigure   = &system{}
	_ resource.ResourceWithImportState = &system{}
)

// var _ resource.ResourceWithConfigValidators = &system{}
// var _ resource.ResourceWithModifyPlan = &system{}
// var _ resource.ResourceWithUpgradeState = &system{}
// var _ resource.ResourceWithValidateConfig = &system{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &system{}
