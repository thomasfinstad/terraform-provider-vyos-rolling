// Package globalservicepppoeserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicepppoeserver

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &servicePppoeServer{}
	_ resource.ResourceWithConfigure = &servicePppoeServer{}
)

// var _ resource.ResourceWithConfigValidators = &servicePppoeServer{}
// var _ resource.ResourceWithModifyPlan = &servicePppoeServer{}
// var _ resource.ResourceWithUpgradeState = &servicePppoeServer{}
// var _ resource.ResourceWithValidateConfig = &servicePppoeServer{}
// var _ resource.ResourceWithImportState = &servicePppoeServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &servicePppoeServer{}
