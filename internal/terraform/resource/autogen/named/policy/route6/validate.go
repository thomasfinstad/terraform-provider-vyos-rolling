// Package namedpolicyroutesix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyroutesix

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &policyRoutesix{}
	_ resource.ResourceWithConfigure = &policyRoutesix{}
)

// var _ resource.ResourceWithConfigValidators = &policyRoutesix{}
// var _ resource.ResourceWithModifyPlan = &policyRoutesix{}
// var _ resource.ResourceWithUpgradeState = &policyRoutesix{}
// var _ resource.ResourceWithValidateConfig = &policyRoutesix{}
// var _ resource.ResourceWithImportState = &policyRoutesix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyRoutesix{}
