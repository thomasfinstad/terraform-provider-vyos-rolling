// Package globalsystemlogin code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemlogin

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemLogin{}
	_ resource.ResourceWithConfigure = &systemLogin{}
)

// var _ resource.ResourceWithConfigValidators = &systemLogin{}
// var _ resource.ResourceWithModifyPlan = &systemLogin{}
// var _ resource.ResourceWithUpgradeState = &systemLogin{}
// var _ resource.ResourceWithValidateConfig = &systemLogin{}
// var _ resource.ResourceWithImportState = &systemLogin{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemLogin{}
