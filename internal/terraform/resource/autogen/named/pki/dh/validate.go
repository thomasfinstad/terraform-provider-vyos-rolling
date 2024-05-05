// Package namedpkidh code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpkidh

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &pkiDh{}
	_ resource.ResourceWithConfigure = &pkiDh{}
)

// var _ resource.ResourceWithConfigValidators = &pkiDh{}
// var _ resource.ResourceWithModifyPlan = &pkiDh{}
// var _ resource.ResourceWithUpgradeState = &pkiDh{}
// var _ resource.ResourceWithValidateConfig = &pkiDh{}
// var _ resource.ResourceWithImportState = &pkiDh{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &pkiDh{}
