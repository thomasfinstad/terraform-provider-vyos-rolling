// Package namedfirewallipvfourpreroutingrawrule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvfourpreroutingrawrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallIPvfourPreroutingRawRule{}
	_ resource.ResourceWithConfigure = &firewallIPvfourPreroutingRawRule{}
)

// var _ resource.ResourceWithConfigValidators = &firewallIPvfourPreroutingRawRule{}
// var _ resource.ResourceWithModifyPlan = &firewallIPvfourPreroutingRawRule{}
// var _ resource.ResourceWithUpgradeState = &firewallIPvfourPreroutingRawRule{}
// var _ resource.ResourceWithValidateConfig = &firewallIPvfourPreroutingRawRule{}
// var _ resource.ResourceWithImportState = &firewallIPvfourPreroutingRawRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallIPvfourPreroutingRawRule{}
