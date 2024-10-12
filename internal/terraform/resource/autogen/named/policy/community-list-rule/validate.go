// Package namedpolicycommunitylistrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicycommunitylistrule

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &policyCommunityListRule{}
	_ resource.ResourceWithConfigure = &policyCommunityListRule{}
)

// var _ resource.ResourceWithConfigValidators = &policyCommunityListRule{}
// var _ resource.ResourceWithModifyPlan = &policyCommunityListRule{}
// var _ resource.ResourceWithUpgradeState = &policyCommunityListRule{}
// var _ resource.ResourceWithValidateConfig = &policyCommunityListRule{}
// var _ resource.ResourceWithImportState = &policyCommunityListRule{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyCommunityListRule{}
