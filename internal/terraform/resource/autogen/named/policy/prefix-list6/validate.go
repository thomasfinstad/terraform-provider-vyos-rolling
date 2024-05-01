// Package namedpolicyprefixlistsix code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpolicyprefixlistsix

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &policyPrefixListsix{}
	_ resource.ResourceWithConfigure = &policyPrefixListsix{}
)

// var _ resource.ResourceWithConfigValidators = &policyPrefixListsix{}
// var _ resource.ResourceWithModifyPlan = &policyPrefixListsix{}
// var _ resource.ResourceWithUpgradeState = &policyPrefixListsix{}
// var _ resource.ResourceWithValidateConfig = &policyPrefixListsix{}
// var _ resource.ResourceWithImportState = &policyPrefixListsix{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &policyPrefixListsix{}
