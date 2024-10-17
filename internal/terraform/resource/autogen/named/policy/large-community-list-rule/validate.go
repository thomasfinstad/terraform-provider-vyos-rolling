// Package namedpolicylargecommunitylistrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicylargecommunitylistrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &policyLargeCommunityListRule{}
	_ resource.ResourceWithConfigure   = &policyLargeCommunityListRule{}
	_ resource.ResourceWithImportState = &policyLargeCommunityListRule{}
)

// var _ resource.ResourceWithConfigValidators = &policyLargeCommunityListRule{}
// var _ resource.ResourceWithModifyPlan = &policyLargeCommunityListRule{}
// var _ resource.ResourceWithUpgradeState = &policyLargeCommunityListRule{}
// var _ resource.ResourceWithValidateConfig = &policyLargeCommunityListRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyLargeCommunityListRule{}
