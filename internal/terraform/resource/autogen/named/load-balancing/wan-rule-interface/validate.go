// Package namedloadbalancingwanruleinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancingwanruleinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &loadBalancingWanRuleInterface{}
	_ resource.ResourceWithConfigure = &loadBalancingWanRuleInterface{}
)

// var _ resource.ResourceWithConfigValidators = &loadBalancingWanRuleInterface{}
// var _ resource.ResourceWithModifyPlan = &loadBalancingWanRuleInterface{}
// var _ resource.ResourceWithUpgradeState = &loadBalancingWanRuleInterface{}
// var _ resource.ResourceWithValidateConfig = &loadBalancingWanRuleInterface{}
// var _ resource.ResourceWithImportState = &loadBalancingWanRuleInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &loadBalancingWanRuleInterface{}