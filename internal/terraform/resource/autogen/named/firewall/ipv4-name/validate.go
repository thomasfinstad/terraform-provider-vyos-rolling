// Package namedfirewallipvfourname code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvfourname

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallIPvfourName{}
	_ resource.ResourceWithConfigure   = &firewallIPvfourName{}
	_ resource.ResourceWithImportState = &firewallIPvfourName{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvfourName{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvfourName{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvfourName{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvfourName{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvfourName{}
