/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedsystemloginradiusserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemloginradiusserver

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemLoginRadiusServer{}
	_ resource.ResourceWithConfigure   = &systemLoginRadiusServer{}
	_ resource.ResourceWithImportState = &systemLoginRadiusServer{}
)

// var _ resource.ResourceWithConfigValidators = &systemLoginRadiusServer{}
// var _ resource.ResourceWithModifyPlan = &systemLoginRadiusServer{}
// var _ resource.ResourceWithUpgradeState = &systemLoginRadiusServer{}
// var _ resource.ResourceWithValidateConfig = &systemLoginRadiusServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemLoginRadiusServer{}
