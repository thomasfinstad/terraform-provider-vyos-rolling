// Package globalservicehttpsapi code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicehttpsapi

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceHTTPSAPI{}
	_ resource.ResourceWithConfigure = &serviceHTTPSAPI{}
)

// var _ resource.ResourceWithConfigValidators = &serviceHTTPSAPI{}
// var _ resource.ResourceWithModifyPlan = &serviceHTTPSAPI{}
// var _ resource.ResourceWithUpgradeState = &serviceHTTPSAPI{}
// var _ resource.ResourceWithValidateConfig = &serviceHTTPSAPI{}
// var _ resource.ResourceWithImportState = &serviceHTTPSAPI{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceHTTPSAPI{}