// Package globalsystemconntrack code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemconntrack

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemConntrack{}
	_ resource.ResourceWithConfigure = &systemConntrack{}
)

// var _ resource.ResourceWithConfigValidators = &systemConntrack{}
// var _ resource.ResourceWithModifyPlan = &systemConntrack{}
// var _ resource.ResourceWithUpgradeState = &systemConntrack{}
// var _ resource.ResourceWithValidateConfig = &systemConntrack{}
// var _ resource.ResourceWithImportState = &systemConntrack{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemConntrack{}
