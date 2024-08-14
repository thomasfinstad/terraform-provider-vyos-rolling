// Package namedinterfacespseudoethernetdhcpvsixoptionspdinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacespseudoethernetdhcpvsixoptionspdinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithConfigure = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithModifyPlan = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithUpgradeState = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithValidateConfig = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithImportState = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{}