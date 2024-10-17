// Package globalservicepppoeservershaper code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicepppoeservershaper

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &servicePppoeServerShaper{}
	_ resource.ResourceWithConfigure   = &servicePppoeServerShaper{}
	_ resource.ResourceWithImportState = &servicePppoeServerShaper{}
)

// var _ resource.ResourceWithConfigValidators = &servicePppoeServerShaper{}
// var _ resource.ResourceWithModifyPlan = &servicePppoeServerShaper{}
// var _ resource.ResourceWithUpgradeState = &servicePppoeServerShaper{}
// var _ resource.ResourceWithValidateConfig = &servicePppoeServerShaper{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &servicePppoeServerShaper{}
