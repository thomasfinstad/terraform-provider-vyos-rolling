// Package namedservicesnmpvthreegroup code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicesnmpvthreegroup

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceSnmpVthreeGroup{}
	_ resource.ResourceWithConfigure   = &serviceSnmpVthreeGroup{}
	_ resource.ResourceWithImportState = &serviceSnmpVthreeGroup{}
)

// var _ resource.ResourceWithConfigValidators = &serviceSnmpVthreeGroup{}
// var _ resource.ResourceWithModifyPlan = &serviceSnmpVthreeGroup{}
// var _ resource.ResourceWithUpgradeState = &serviceSnmpVthreeGroup{}
// var _ resource.ResourceWithValidateConfig = &serviceSnmpVthreeGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceSnmpVthreeGroup{}
