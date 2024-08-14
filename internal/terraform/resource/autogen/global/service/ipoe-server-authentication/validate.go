// Package globalserviceipoeserverauthentication code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceipoeserverauthentication

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceIPoeServerAuthentication{}
	_ resource.ResourceWithConfigure = &serviceIPoeServerAuthentication{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIPoeServerAuthentication{}
// var _ resource.ResourceWithModifyPlan = &serviceIPoeServerAuthentication{}
// var _ resource.ResourceWithUpgradeState = &serviceIPoeServerAuthentication{}
// var _ resource.ResourceWithValidateConfig = &serviceIPoeServerAuthentication{}
// var _ resource.ResourceWithImportState = &serviceIPoeServerAuthentication{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIPoeServerAuthentication{}