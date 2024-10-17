// Package namedpolicyextcommunitylist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyextcommunitylist

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &policyExtcommunityList{}
	_ resource.ResourceWithConfigure   = &policyExtcommunityList{}
	_ resource.ResourceWithImportState = &policyExtcommunityList{}
)

// var _ resource.ResourceWithConfigValidators = &policyExtcommunityList{}
// var _ resource.ResourceWithModifyPlan = &policyExtcommunityList{}
// var _ resource.ResourceWithUpgradeState = &policyExtcommunityList{}
// var _ resource.ResourceWithValidateConfig = &policyExtcommunityList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyExtcommunityList{}
