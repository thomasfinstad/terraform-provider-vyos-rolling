// Package namedsystemloginuser code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemloginuser

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemLoginUser{}
	_ resource.ResourceWithConfigure   = &systemLoginUser{}
	_ resource.ResourceWithImportState = &systemLoginUser{}
)

// var _ resource.ResourceWithConfigValidators = &systemLoginUser{}
// var _ resource.ResourceWithModifyPlan = &systemLoginUser{}
// var _ resource.ResourceWithUpgradeState = &systemLoginUser{}
// var _ resource.ResourceWithValidateConfig = &systemLoginUser{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemLoginUser{}
