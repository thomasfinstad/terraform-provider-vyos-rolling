// Package namedsystemconntrackignoreipvsixrule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemconntrackignoreipvsixrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemConntrackIgnoreIPvsixRule{}
	_ resource.ResourceWithConfigure = &systemConntrackIgnoreIPvsixRule{}
)

// var _ resource.ResourceWithConfigValidators = &systemConntrackIgnoreIPvsixRule{}
// var _ resource.ResourceWithModifyPlan = &systemConntrackIgnoreIPvsixRule{}
// var _ resource.ResourceWithUpgradeState = &systemConntrackIgnoreIPvsixRule{}
// var _ resource.ResourceWithValidateConfig = &systemConntrackIgnoreIPvsixRule{}
// var _ resource.ResourceWithImportState = &systemConntrackIgnoreIPvsixRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemConntrackIgnoreIPvsixRule{}
