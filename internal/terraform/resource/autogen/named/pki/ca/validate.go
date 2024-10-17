// Package namedpkica code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpkica

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &pkiCa{}
	_ resource.ResourceWithConfigure   = &pkiCa{}
	_ resource.ResourceWithImportState = &pkiCa{}
)

// var _ resource.ResourceWithConfigValidators = &pkiCa{}
// var _ resource.ResourceWithModifyPlan = &pkiCa{}
// var _ resource.ResourceWithUpgradeState = &pkiCa{}
// var _ resource.ResourceWithValidateConfig = &pkiCa{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &pkiCa{}
