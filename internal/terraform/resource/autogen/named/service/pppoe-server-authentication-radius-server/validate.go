// Package namedservicepppoeserverauthenticationradiusserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicepppoeserverauthenticationradiusserver

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &servicePppoeServerAuthenticationRadiusServer{}
	_ resource.ResourceWithConfigure = &servicePppoeServerAuthenticationRadiusServer{}
)

// var _ resource.ResourceWithConfigValidators = &servicePppoeServerAuthenticationRadiusServer{}
// var _ resource.ResourceWithModifyPlan = &servicePppoeServerAuthenticationRadiusServer{}
// var _ resource.ResourceWithUpgradeState = &servicePppoeServerAuthenticationRadiusServer{}
// var _ resource.ResourceWithValidateConfig = &servicePppoeServerAuthenticationRadiusServer{}
// var _ resource.ResourceWithImportState = &servicePppoeServerAuthenticationRadiusServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &servicePppoeServerAuthenticationRadiusServer{}
