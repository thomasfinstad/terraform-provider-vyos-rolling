// Package globalsystemoptionhttpclient code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemoptionhttpclient

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemOptionHTTPClient{}
	_ resource.ResourceWithConfigure = &systemOptionHTTPClient{}
)

// var _ resource.ResourceWithConfigValidators = &systemOptionHTTPClient{}
// var _ resource.ResourceWithModifyPlan = &systemOptionHTTPClient{}
// var _ resource.ResourceWithUpgradeState = &systemOptionHTTPClient{}
// var _ resource.ResourceWithValidateConfig = &systemOptionHTTPClient{}
// var _ resource.ResourceWithImportState = &systemOptionHTTPClient{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemOptionHTTPClient{}