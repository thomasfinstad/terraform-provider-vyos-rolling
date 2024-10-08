// Package namedsystemconntrackignoreipvfourrule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemconntrackignoreipvfourrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemConntrackIgnoreIPvfourRule{}
	_ resource.ResourceWithConfigure = &systemConntrackIgnoreIPvfourRule{}
)

// var _ resource.ResourceWithConfigValidators = &systemConntrackIgnoreIPvfourRule{}
// var _ resource.ResourceWithModifyPlan = &systemConntrackIgnoreIPvfourRule{}
// var _ resource.ResourceWithUpgradeState = &systemConntrackIgnoreIPvfourRule{}
// var _ resource.ResourceWithValidateConfig = &systemConntrackIgnoreIPvfourRule{}
// var _ resource.ResourceWithImportState = &systemConntrackIgnoreIPvfourRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemConntrackIgnoreIPvfourRule{}
