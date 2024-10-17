// Package namedinterfacesopenvpn code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesopenvpn

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesOpenvpn{}
	_ resource.ResourceWithConfigure   = &interfacesOpenvpn{}
	_ resource.ResourceWithImportState = &interfacesOpenvpn{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesOpenvpn{}
// var _ resource.ResourceWithModifyPlan = &interfacesOpenvpn{}
// var _ resource.ResourceWithUpgradeState = &interfacesOpenvpn{}
// var _ resource.ResourceWithValidateConfig = &interfacesOpenvpn{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesOpenvpn{}
