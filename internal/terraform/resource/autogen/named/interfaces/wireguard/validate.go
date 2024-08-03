// Package namedinterfaceswireguard code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswireguard

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesWireguard{}
	_ resource.ResourceWithConfigure = &interfacesWireguard{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesWireguard{}
// var _ resource.ResourceWithModifyPlan = &interfacesWireguard{}
// var _ resource.ResourceWithUpgradeState = &interfacesWireguard{}
// var _ resource.ResourceWithValidateConfig = &interfacesWireguard{}
// var _ resource.ResourceWithImportState = &interfacesWireguard{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesWireguard{}
