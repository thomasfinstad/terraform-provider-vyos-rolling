// Package namedcontainernamevolume code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedcontainernamevolume

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &containerNameVolume{}
	_ resource.ResourceWithConfigure = &containerNameVolume{}
)

// var _ resource.ResourceWithConfigValidators = &containerNameVolume{}
// var _ resource.ResourceWithModifyPlan = &containerNameVolume{}
// var _ resource.ResourceWithUpgradeState = &containerNameVolume{}
// var _ resource.ResourceWithValidateConfig = &containerNameVolume{}
// var _ resource.ResourceWithImportState = &containerNameVolume{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &containerNameVolume{}
