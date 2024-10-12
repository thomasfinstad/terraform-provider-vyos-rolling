// Package globalsystemloginradius code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemloginradius

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemLoginRadius{}
	_ resource.ResourceWithConfigure = &systemLoginRadius{}
)

// var _ resource.ResourceWithConfigValidators = &systemLoginRadius{}
// var _ resource.ResourceWithModifyPlan = &systemLoginRadius{}
// var _ resource.ResourceWithUpgradeState = &systemLoginRadius{}
// var _ resource.ResourceWithValidateConfig = &systemLoginRadius{}
// var _ resource.ResourceWithImportState = &systemLoginRadius{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemLoginRadius{}
