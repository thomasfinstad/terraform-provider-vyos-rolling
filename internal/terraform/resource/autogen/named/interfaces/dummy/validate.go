// Package namedinterfacesdummy code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesdummy

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesDummy{}
	_ resource.ResourceWithConfigure = &interfacesDummy{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesDummy{}
// var _ resource.ResourceWithModifyPlan = &interfacesDummy{}
// var _ resource.ResourceWithUpgradeState = &interfacesDummy{}
// var _ resource.ResourceWithValidateConfig = &interfacesDummy{}
// var _ resource.ResourceWithImportState = &interfacesDummy{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesDummy{}