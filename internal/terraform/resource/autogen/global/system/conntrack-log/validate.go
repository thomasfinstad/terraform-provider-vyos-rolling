// Package globalsystemconntracklog code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemconntracklog

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &systemConntrackLog{}
	_ resource.ResourceWithConfigure   = &systemConntrackLog{}
	_ resource.ResourceWithImportState = &systemConntrackLog{}
)

// var _ resource.ResourceWithConfigValidators = &systemConntrackLog{}
// var _ resource.ResourceWithModifyPlan = &systemConntrackLog{}
// var _ resource.ResourceWithUpgradeState = &systemConntrackLog{}
// var _ resource.ResourceWithValidateConfig = &systemConntrackLog{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemConntrackLog{}
