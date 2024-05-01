// Package globalvrf code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalvrf

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrf{}
	_ resource.ResourceWithConfigure = &vrf{}
)

// var _ resource.ResourceWithConfigValidators = &vrf{}
// var _ resource.ResourceWithModifyPlan = &vrf{}
// var _ resource.ResourceWithUpgradeState = &vrf{}
// var _ resource.ResourceWithValidateConfig = &vrf{}
// var _ resource.ResourceWithImportState = &vrf{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrf{}
