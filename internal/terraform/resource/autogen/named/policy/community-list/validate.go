// Package namedpolicycommunitylist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicycommunitylist

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &policyCommunityList{}
	_ resource.ResourceWithConfigure   = &policyCommunityList{}
	_ resource.ResourceWithImportState = &policyCommunityList{}
)

// var _ resource.ResourceWithConfigValidators = &policyCommunityList{}
// var _ resource.ResourceWithModifyPlan = &policyCommunityList{}
// var _ resource.ResourceWithUpgradeState = &policyCommunityList{}
// var _ resource.ResourceWithValidateConfig = &policyCommunityList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyCommunityList{}
