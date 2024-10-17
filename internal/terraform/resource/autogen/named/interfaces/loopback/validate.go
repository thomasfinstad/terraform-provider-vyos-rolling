// Package namedinterfacesloopback code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesloopback

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesLoopback{}
	_ resource.ResourceWithConfigure   = &interfacesLoopback{}
	_ resource.ResourceWithImportState = &interfacesLoopback{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesLoopback{}
// var _ resource.ResourceWithModifyPlan = &interfacesLoopback{}
// var _ resource.ResourceWithUpgradeState = &interfacesLoopback{}
// var _ resource.ResourceWithValidateConfig = &interfacesLoopback{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesLoopback{}
