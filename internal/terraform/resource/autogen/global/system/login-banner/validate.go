// Package globalsystemloginbanner code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemloginbanner

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemLoginBanner{}
	_ resource.ResourceWithConfigure   = &systemLoginBanner{}
	_ resource.ResourceWithImportState = &systemLoginBanner{}
)

// var _ resource.ResourceWithConfigValidators = &systemLoginBanner{}
// var _ resource.ResourceWithModifyPlan = &systemLoginBanner{}
// var _ resource.ResourceWithUpgradeState = &systemLoginBanner{}
// var _ resource.ResourceWithValidateConfig = &systemLoginBanner{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemLoginBanner{}
