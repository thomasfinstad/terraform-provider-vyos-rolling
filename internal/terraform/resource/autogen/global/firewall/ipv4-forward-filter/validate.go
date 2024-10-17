// Package globalfirewallipvfourforwardfilter code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallipvfourforwardfilter

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallIPvfourForwardFilter{}
	_ resource.ResourceWithConfigure   = &firewallIPvfourForwardFilter{}
	_ resource.ResourceWithImportState = &firewallIPvfourForwardFilter{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvfourForwardFilter{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvfourForwardFilter{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvfourForwardFilter{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvfourForwardFilter{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvfourForwardFilter{}
