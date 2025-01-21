/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfacesvirtualethernetvifsvifcdhcpvsixoptionspd code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesvirtualethernetvifsvifcdhcpvsixoptionspd

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (pd) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd{}
	_ resource.ResourceWithConfigure   = &interfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd{}
	_ resource.ResourceWithImportState = &interfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd{}
// var _ resource.ResourceWithModifyPlan = &interfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd{}
// var _ resource.ResourceWithUpgradeState = &interfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd{}
// var _ resource.ResourceWithValidateConfig = &interfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd{}
