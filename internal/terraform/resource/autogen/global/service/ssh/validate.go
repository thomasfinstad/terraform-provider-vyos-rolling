// Package globalservicetcp code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicetcp

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceTCP{}
	_ resource.ResourceWithConfigure = &serviceTCP{}
)

// var _ resource.ResourceWithConfigValidators = &serviceTCP{}
// var _ resource.ResourceWithModifyPlan = &serviceTCP{}
// var _ resource.ResourceWithUpgradeState = &serviceTCP{}
// var _ resource.ResourceWithValidateConfig = &serviceTCP{}
// var _ resource.ResourceWithImportState = &serviceTCP{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceTCP{}
