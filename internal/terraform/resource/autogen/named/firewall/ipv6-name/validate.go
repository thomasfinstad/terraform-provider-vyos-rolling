// Package namedfirewallipvsixname code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvsixname

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallIPvsixName{}
	_ resource.ResourceWithConfigure = &firewallIPvsixName{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvsixName{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvsixName{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvsixName{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvsixName{}
// var _ resource.ResourceWithImportState = &firewallIPvsixName{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvsixName{}
