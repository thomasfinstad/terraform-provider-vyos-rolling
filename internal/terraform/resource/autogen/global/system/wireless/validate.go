/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalsystemwireless code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemwireless

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemWireless{}
	_ resource.ResourceWithConfigure   = &systemWireless{}
	_ resource.ResourceWithImportState = &systemWireless{}
)

// var _ resource.ResourceWithConfigValidators = &systemWireless{}
// var _ resource.ResourceWithModifyPlan = &systemWireless{}
// var _ resource.ResourceWithUpgradeState = &systemWireless{}
// var _ resource.ResourceWithValidateConfig = &systemWireless{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemWireless{}
