// Package namedinterfacesethernet code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesethernet

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesEthernet{}
	_ resource.ResourceWithConfigure   = &interfacesEthernet{}
	_ resource.ResourceWithImportState = &interfacesEthernet{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesEthernet{}
// var _ resource.ResourceWithModifyPlan = &interfacesEthernet{}
// var _ resource.ResourceWithUpgradeState = &interfacesEthernet{}
// var _ resource.ResourceWithValidateConfig = &interfacesEthernet{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesEthernet{}
