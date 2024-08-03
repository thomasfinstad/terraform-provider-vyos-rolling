// Package globalsystemconntracklogeventnew code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemconntracklogeventnew

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemConntrackLogEventNew{}
	_ resource.ResourceWithConfigure = &systemConntrackLogEventNew{}
)

// var _ resource.ResourceWithConfigValidators = &systemConntrackLogEventNew{}
// var _ resource.ResourceWithModifyPlan = &systemConntrackLogEventNew{}
// var _ resource.ResourceWithUpgradeState = &systemConntrackLogEventNew{}
// var _ resource.ResourceWithValidateConfig = &systemConntrackLogEventNew{}
// var _ resource.ResourceWithImportState = &systemConntrackLogEventNew{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemConntrackLogEventNew{}
