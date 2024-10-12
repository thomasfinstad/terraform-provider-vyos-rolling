// Package globalsystemipvsixnht code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemipvsixnht

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemIPvsixNht{}
	_ resource.ResourceWithConfigure = &systemIPvsixNht{}
)

// var _ resource.ResourceWithConfigValidators = &systemIPvsixNht{}
// var _ resource.ResourceWithModifyPlan = &systemIPvsixNht{}
// var _ resource.ResourceWithUpgradeState = &systemIPvsixNht{}
// var _ resource.ResourceWithValidateConfig = &systemIPvsixNht{}
// var _ resource.ResourceWithImportState = &systemIPvsixNht{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemIPvsixNht{}
