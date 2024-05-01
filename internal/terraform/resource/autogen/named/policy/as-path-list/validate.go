// Package namedpolicyaspathlist code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyaspathlist

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &policyAsPathList{}
	_ resource.ResourceWithConfigure = &policyAsPathList{}
)

// var _ resource.ResourceWithConfigValidators = &policyAsPathList{}
// var _ resource.ResourceWithModifyPlan = &policyAsPathList{}
// var _ resource.ResourceWithUpgradeState = &policyAsPathList{}
// var _ resource.ResourceWithValidateConfig = &policyAsPathList{}
// var _ resource.ResourceWithImportState = &policyAsPathList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyAsPathList{}
