/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalsystemoption code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemoption

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemOption{}
	_ resource.ResourceWithConfigure   = &systemOption{}
	_ resource.ResourceWithImportState = &systemOption{}
)

// var _ resource.ResourceWithConfigValidators = &systemOption{}
// var _ resource.ResourceWithModifyPlan = &systemOption{}
// var _ resource.ResourceWithUpgradeState = &systemOption{}
// var _ resource.ResourceWithValidateConfig = &systemOption{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemOption{}
