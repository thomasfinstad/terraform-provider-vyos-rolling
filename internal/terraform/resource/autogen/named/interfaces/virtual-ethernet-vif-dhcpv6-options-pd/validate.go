/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfacesvirtualethernetvifdhcpvsixoptionspd code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesvirtualethernetvifdhcpvsixoptionspd

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (pd) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesVirtualEthernetVifDhcpvsixOptionsPd{}
	_ resource.ResourceWithConfigure   = &interfacesVirtualEthernetVifDhcpvsixOptionsPd{}
	_ resource.ResourceWithImportState = &interfacesVirtualEthernetVifDhcpvsixOptionsPd{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesVirtualEthernetVifDhcpvsixOptionsPd{}
// var _ resource.ResourceWithModifyPlan = &interfacesVirtualEthernetVifDhcpvsixOptionsPd{}
// var _ resource.ResourceWithUpgradeState = &interfacesVirtualEthernetVifDhcpvsixOptionsPd{}
// var _ resource.ResourceWithValidateConfig = &interfacesVirtualEthernetVifDhcpvsixOptionsPd{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesVirtualEthernetVifDhcpvsixOptionsPd{}
