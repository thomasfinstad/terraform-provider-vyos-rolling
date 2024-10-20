/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedinterfacespseudoethernetdhcpvsixoptionspd code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacespseudoethernetdhcpvsixoptionspd

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesPseudoEthernetDhcpvsixOptionsPd{}
	_ resource.ResourceWithConfigure   = &interfacesPseudoEthernetDhcpvsixOptionsPd{}
	_ resource.ResourceWithImportState = &interfacesPseudoEthernetDhcpvsixOptionsPd{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesPseudoEthernetDhcpvsixOptionsPd{}
// var _ resource.ResourceWithModifyPlan = &interfacesPseudoEthernetDhcpvsixOptionsPd{}
// var _ resource.ResourceWithUpgradeState = &interfacesPseudoEthernetDhcpvsixOptionsPd{}
// var _ resource.ResourceWithValidateConfig = &interfacesPseudoEthernetDhcpvsixOptionsPd{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesPseudoEthernetDhcpvsixOptionsPd{}
