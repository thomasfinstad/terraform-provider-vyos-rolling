// Package globalservicesaltminion code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicesaltminion

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceSaltMinion{}
	_ resource.ResourceWithConfigure = &serviceSaltMinion{}
)

// var _ resource.ResourceWithConfigValidators = &serviceSaltMinion{}
// var _ resource.ResourceWithModifyPlan = &serviceSaltMinion{}
// var _ resource.ResourceWithUpgradeState = &serviceSaltMinion{}
// var _ resource.ResourceWithValidateConfig = &serviceSaltMinion{}
// var _ resource.ResourceWithImportState = &serviceSaltMinion{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceSaltMinion{}
