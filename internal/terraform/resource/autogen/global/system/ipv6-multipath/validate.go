// Package globalsystemipvsixmultipath code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemipvsixmultipath

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemIPvsixMultIPath{}
	_ resource.ResourceWithConfigure   = &systemIPvsixMultIPath{}
	_ resource.ResourceWithImportState = &systemIPvsixMultIPath{}
)

// var _ resource.ResourceWithConfigValidators = &systemIPvsixMultIPath{}
// var _ resource.ResourceWithModifyPlan = &systemIPvsixMultIPath{}
// var _ resource.ResourceWithUpgradeState = &systemIPvsixMultIPath{}
// var _ resource.ResourceWithValidateConfig = &systemIPvsixMultIPath{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemIPvsixMultIPath{}
