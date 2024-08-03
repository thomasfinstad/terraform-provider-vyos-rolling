// Package namedinterfacespseudoethernetvif code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacespseudoethernetvif

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesPseudoEthernetVif{}
	_ resource.ResourceWithConfigure = &interfacesPseudoEthernetVif{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesPseudoEthernetVif{}
// var _ resource.ResourceWithModifyPlan = &interfacesPseudoEthernetVif{}
// var _ resource.ResourceWithUpgradeState = &interfacesPseudoEthernetVif{}
// var _ resource.ResourceWithValidateConfig = &interfacesPseudoEthernetVif{}
// var _ resource.ResourceWithImportState = &interfacesPseudoEthernetVif{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesPseudoEthernetVif{}
