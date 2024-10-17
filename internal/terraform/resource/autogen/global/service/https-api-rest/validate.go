// Package globalservicehttpsapirest code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicehttpsapirest

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceHTTPSAPIRest{}
	_ resource.ResourceWithConfigure   = &serviceHTTPSAPIRest{}
	_ resource.ResourceWithImportState = &serviceHTTPSAPIRest{}
)

// var _ resource.ResourceWithConfigValidators = &serviceHTTPSAPIRest{}
// var _ resource.ResourceWithModifyPlan = &serviceHTTPSAPIRest{}
// var _ resource.ResourceWithUpgradeState = &serviceHTTPSAPIRest{}
// var _ resource.ResourceWithValidateConfig = &serviceHTTPSAPIRest{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceHTTPSAPIRest{}
