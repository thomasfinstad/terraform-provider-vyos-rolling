/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedsystemipvsixprotocol code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemipvsixprotocol

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemIPvsixProtocol{}
	_ resource.ResourceWithConfigure   = &systemIPvsixProtocol{}
	_ resource.ResourceWithImportState = &systemIPvsixProtocol{}
)

// var _ resource.ResourceWithConfigValidators = &systemIPvsixProtocol{}
// var _ resource.ResourceWithModifyPlan = &systemIPvsixProtocol{}
// var _ resource.ResourceWithUpgradeState = &systemIPvsixProtocol{}
// var _ resource.ResourceWithValidateConfig = &systemIPvsixProtocol{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemIPvsixProtocol{}
