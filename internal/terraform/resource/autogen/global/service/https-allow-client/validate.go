// Package globalservicehttpsallowclient code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicehttpsallowclient

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceHTTPSAllowClient{}
	_ resource.ResourceWithConfigure = &serviceHTTPSAllowClient{}
)

// var _ resource.ResourceWithConfigValidators = &serviceHTTPSAllowClient{}
// var _ resource.ResourceWithModifyPlan = &serviceHTTPSAllowClient{}
// var _ resource.ResourceWithUpgradeState = &serviceHTTPSAllowClient{}
// var _ resource.ResourceWithValidateConfig = &serviceHTTPSAllowClient{}
// var _ resource.ResourceWithImportState = &serviceHTTPSAllowClient{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceHTTPSAllowClient{}