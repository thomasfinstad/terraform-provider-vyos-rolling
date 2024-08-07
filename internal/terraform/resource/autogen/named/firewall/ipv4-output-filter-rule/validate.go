// Package namedfirewallipvfouroutputfilterrule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvfouroutputfilterrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallIPvfourOutputFilterRule{}
	_ resource.ResourceWithConfigure = &firewallIPvfourOutputFilterRule{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvfourOutputFilterRule{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvfourOutputFilterRule{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvfourOutputFilterRule{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvfourOutputFilterRule{}
// var _ resource.ResourceWithImportState = &firewallIPvfourOutputFilterRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvfourOutputFilterRule{}
