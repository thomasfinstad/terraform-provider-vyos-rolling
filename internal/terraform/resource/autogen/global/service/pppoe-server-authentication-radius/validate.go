// Package globalservicepppoeserverauthenticationradius code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicepppoeserverauthenticationradius

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &servicePppoeServerAuthenticationRadius{}
	_ resource.ResourceWithConfigure = &servicePppoeServerAuthenticationRadius{}
)

// var _ resource.ResourceWithConfigValidators = &servicePppoeServerAuthenticationRadius{}
// var _ resource.ResourceWithModifyPlan = &servicePppoeServerAuthenticationRadius{}
// var _ resource.ResourceWithUpgradeState = &servicePppoeServerAuthenticationRadius{}
// var _ resource.ResourceWithValidateConfig = &servicePppoeServerAuthenticationRadius{}
// var _ resource.ResourceWithImportState = &servicePppoeServerAuthenticationRadius{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &servicePppoeServerAuthenticationRadius{}
