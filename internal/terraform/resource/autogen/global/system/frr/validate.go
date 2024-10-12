// Package globalsystemfrr code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemfrr

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemFrr{}
	_ resource.ResourceWithConfigure = &systemFrr{}
)

// var _ resource.ResourceWithConfigValidators = &systemFrr{}
// var _ resource.ResourceWithModifyPlan = &systemFrr{}
// var _ resource.ResourceWithUpgradeState = &systemFrr{}
// var _ resource.ResourceWithValidateConfig = &systemFrr{}
// var _ resource.ResourceWithImportState = &systemFrr{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemFrr{}
