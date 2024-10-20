/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalsystemconsole code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemconsole

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemConsole{}
	_ resource.ResourceWithConfigure   = &systemConsole{}
	_ resource.ResourceWithImportState = &systemConsole{}
)

// var _ resource.ResourceWithConfigValidators = &systemConsole{}
// var _ resource.ResourceWithModifyPlan = &systemConsole{}
// var _ resource.ResourceWithUpgradeState = &systemConsole{}
// var _ resource.ResourceWithValidateConfig = &systemConsole{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemConsole{}
