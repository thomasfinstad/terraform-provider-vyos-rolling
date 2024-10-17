// Package namedinterfacesbridge code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesbridge

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesBrIDge{}
	_ resource.ResourceWithConfigure   = &interfacesBrIDge{}
	_ resource.ResourceWithImportState = &interfacesBrIDge{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesBrIDge{}
// var _ resource.ResourceWithModifyPlan = &interfacesBrIDge{}
// var _ resource.ResourceWithUpgradeState = &interfacesBrIDge{}
// var _ resource.ResourceWithValidateConfig = &interfacesBrIDge{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesBrIDge{}
