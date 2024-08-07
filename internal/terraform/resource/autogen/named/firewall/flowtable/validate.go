// Package namedfirewallflowtable code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallflowtable

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &firewallFlowtable{}
	_ resource.ResourceWithConfigure = &firewallFlowtable{}
)

// var _ resource.ResourceWithConfigValidators = &firewallFlowtable{}
// var _ resource.ResourceWithModifyPlan = &firewallFlowtable{}
// var _ resource.ResourceWithUpgradeState = &firewallFlowtable{}
// var _ resource.ResourceWithValidateConfig = &firewallFlowtable{}
// var _ resource.ResourceWithImportState = &firewallFlowtable{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallFlowtable{}
