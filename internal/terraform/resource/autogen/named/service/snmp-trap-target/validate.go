// Package namedservicesnmptraptarget code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicesnmptraptarget

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceSnmpTrapTarget{}
	_ resource.ResourceWithConfigure   = &serviceSnmpTrapTarget{}
	_ resource.ResourceWithImportState = &serviceSnmpTrapTarget{}
)

// var _ resource.ResourceWithConfigValidators = &serviceSnmpTrapTarget{}
// var _ resource.ResourceWithModifyPlan = &serviceSnmpTrapTarget{}
// var _ resource.ResourceWithUpgradeState = &serviceSnmpTrapTarget{}
// var _ resource.ResourceWithValidateConfig = &serviceSnmpTrapTarget{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceSnmpTrapTarget{}
