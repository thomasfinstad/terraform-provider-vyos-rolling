// Package namedfirewallipvfourinputfilterrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvfourinputfilterrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallIPvfourInputFilterRule{}
	_ resource.ResourceWithConfigure = &firewallIPvfourInputFilterRule{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvfourInputFilterRule{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvfourInputFilterRule{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvfourInputFilterRule{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvfourInputFilterRule{}
// var _ resource.ResourceWithImportState = &firewallIPvfourInputFilterRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvfourInputFilterRule{}
