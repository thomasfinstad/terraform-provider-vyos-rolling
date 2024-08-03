// Package namedsystemconsoledevice code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemconsoledevice

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemConsoleDevice{}
	_ resource.ResourceWithConfigure = &systemConsoleDevice{}
)

// var _ resource.ResourceWithConfigValidators = &systemConsoleDevice{}
// var _ resource.ResourceWithModifyPlan = &systemConsoleDevice{}
// var _ resource.ResourceWithUpgradeState = &systemConsoleDevice{}
// var _ resource.ResourceWithValidateConfig = &systemConsoleDevice{}
// var _ resource.ResourceWithImportState = &systemConsoleDevice{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemConsoleDevice{}
