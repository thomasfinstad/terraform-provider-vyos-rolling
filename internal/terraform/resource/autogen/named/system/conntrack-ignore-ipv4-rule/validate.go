/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedsystemconntrackignoreipvfourrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemconntrackignoreipvfourrule

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemConntrackIgnoreIPvfourRule{}
	_ resource.ResourceWithConfigure   = &systemConntrackIgnoreIPvfourRule{}
	_ resource.ResourceWithImportState = &systemConntrackIgnoreIPvfourRule{}
)

// var _ resource.ResourceWithConfigValidators = &systemConntrackIgnoreIPvfourRule{}
// var _ resource.ResourceWithModifyPlan = &systemConntrackIgnoreIPvfourRule{}
// var _ resource.ResourceWithUpgradeState = &systemConntrackIgnoreIPvfourRule{}
// var _ resource.ResourceWithValidateConfig = &systemConntrackIgnoreIPvfourRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemConntrackIgnoreIPvfourRule{}
