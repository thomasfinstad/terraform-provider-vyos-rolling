// Package globalserviceslaowampserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceslaowampserver

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceSLAOwampServer{}
	_ resource.ResourceWithConfigure   = &serviceSLAOwampServer{}
	_ resource.ResourceWithImportState = &serviceSLAOwampServer{}
)

// var _ resource.ResourceWithConfigValidators = &serviceSLAOwampServer{}
// var _ resource.ResourceWithModifyPlan = &serviceSLAOwampServer{}
// var _ resource.ResourceWithUpgradeState = &serviceSLAOwampServer{}
// var _ resource.ResourceWithValidateConfig = &serviceSLAOwampServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceSLAOwampServer{}
