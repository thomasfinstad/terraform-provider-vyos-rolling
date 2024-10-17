// Package namedservicewebproxycachepeer code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicewebproxycachepeer

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceWebproxyCachePeer{}
	_ resource.ResourceWithConfigure   = &serviceWebproxyCachePeer{}
	_ resource.ResourceWithImportState = &serviceWebproxyCachePeer{}
)

// var _ resource.ResourceWithConfigValidators = &serviceWebproxyCachePeer{}
// var _ resource.ResourceWithModifyPlan = &serviceWebproxyCachePeer{}
// var _ resource.ResourceWithUpgradeState = &serviceWebproxyCachePeer{}
// var _ resource.ResourceWithValidateConfig = &serviceWebproxyCachePeer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceWebproxyCachePeer{}
