// Package globalserviceipoeserverlimits code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceipoeserverlimits

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceIPoeServerLimits{}
	_ resource.ResourceWithConfigure = &serviceIPoeServerLimits{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIPoeServerLimits{}
// var _ resource.ResourceWithModifyPlan = &serviceIPoeServerLimits{}
// var _ resource.ResourceWithUpgradeState = &serviceIPoeServerLimits{}
// var _ resource.ResourceWithValidateConfig = &serviceIPoeServerLimits{}
// var _ resource.ResourceWithImportState = &serviceIPoeServerLimits{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIPoeServerLimits{}